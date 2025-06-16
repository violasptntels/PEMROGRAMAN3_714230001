package repository

import (
	"context"
	"fmt"
	"inibackend/config"
	"inibackend/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertMahasiswa(ctx context.Context, mhs model.Mahasiswa) (insertedID interface{}, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.MahasiswaCollection)

	// Cek apakah NPM sudah ada
	filter := bson.M{"npm": mhs.NPM}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		fmt.Printf("InsertMahasiswa - Count: %v\n", err)
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("NPM %v sudah terdaftar", mhs.NPM)
	}

	// Insert jika NPM belum ada
	insertResult, err := collection.InsertOne(ctx, mhs)
	if err != nil {
		fmt.Printf("InsertMahasiswa - Insert: %v\n", err)
		return nil, err
	}

	return insertResult.InsertedID, nil
}

func GetMahasiswaByNPM(ctx context.Context, npm int) (mhs *model.Mahasiswa, err error) {
	mahasiswa := config.MongoConnect(config.DBName).Collection(config.MahasiswaCollection)
	filter := bson.M{"npm": npm}
	err = mahasiswa.FindOne(ctx, filter).Decode(&mhs)
	if err != nil {
		//jika data yang ditemukan akan error
		if err == mongo.ErrNoDocuments {
			return nil, nil //kembalikan nilai nil jika data tidak ada
		}
		return nil, fmt.Errorf("terjadi kesalahan mengambil data: %v\n", err)
	}
	return mhs, nil
}

func GetAllMahasiswa(ctx context.Context) ([]model.Mahasiswa, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.MahasiswaCollection)
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		fmt.Println("GetAllMahasiswa (Find):", err)
		return nil, err
	}

	var data []model.Mahasiswa
	if err := cursor.All(ctx, &data); err != nil {
		fmt.Println("GetAllMahasiswa (Decode):", err)
		return nil, err
	}

	return data, nil
}

func UpdateMahasiswa(ctx context.Context, npm int, update model.Mahasiswa) (updatedNPM int, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.MahasiswaCollection)

	filter := bson.M{"npm": npm}
	updateData := bson.M{"$set": update}

	result, err := collection.UpdateOne(ctx, filter, updateData)
	if err != nil {
		fmt.Printf("UpdateMahasiswa: %v\n", err)
		return 0, err
	}
	if result.ModifiedCount == 0 {
		return 0, fmt.Errorf("tidak ada data yang diupdate untuk NPM %v", npm)
	}
	return npm, nil
}

func DeleteMahasiswa(ctx context.Context, npm int) (deletedNPM int, err error) {
	collection := config.MongoConnect(config.DBName).Collection(config.MahasiswaCollection)

	filter := bson.M{"npm": npm}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Printf("DeleteMahasiswa: %v\n", err)
		return 0, err
	}
	if result.DeletedCount == 0 {
		return 0, fmt.Errorf("tidak ada data yang dihapus untuk NPM %v", npm)
	}
	return npm, nil
}

func FindUserByUsername(ctx context.Context, username string) (*model.UserLogin, error) {
	userCollection := config.MongoConnect(config.DBName).Collection(config.UserCollection)

	var user model.UserLogin
	filter := bson.M{"username": username}

	err := userCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("username %s tidak ditemukan", username)
		}
		return nil, err
	}

	return &user, nil
}

func InsertUser(ctx context.Context, user model.UserLogin) (interface{}, error) {
	collection := config.MongoConnect(config.DBName).Collection(config.UserCollection)

	// Cek apakah username sudah ada
	filter := bson.M{"username": user.Username}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, fmt.Errorf("username %s sudah digunakan", user.Username)
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}