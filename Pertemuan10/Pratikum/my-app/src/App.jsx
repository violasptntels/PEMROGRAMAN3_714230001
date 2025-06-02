import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { LayoutDashboard } from "./components/layouts/LayoutDashboard";
import { Dashboard } from "./pages/Dashboard";
import { MahasiswaPage } from "./pages/MahasiswaPage";
import { TambahMahasiswaPage } from "./pages/TambahMahasiswaPage";
import { EditMahasiswaPage } from "./pages/EditMahasiswaPage";

function App() {
  return (
    <Router>
      <LayoutDashboard>
        <Routes>
          <Route path="/" element={<Dashboard />} />
          <Route path="/mahasiswa" element={<MahasiswaPage />} />
          <Route path="/mahasiswa/tambah" element={<TambahMahasiswaPage />} />
          <Route path="/mahasiswa/edit/:npm" element={<EditMahasiswaPage />} />
        </Routes>
      </LayoutDashboard>
    </Router>
  );
}

export default App