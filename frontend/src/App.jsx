import { Routes, Route, Navigate } from "react-router-dom";
import Navbar from "./components/Navbar";
import Gallery from "./pages/Gallery";
import Albums from "./pages/Albums";
import Vault from "./pages/Vault";
import Login from "./pages/Login";
import Register from "./pages/Register";

function App() {
  return (
    <div>
      <Navbar />
      <div style={{ padding: "12px" }}>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route
            path="/"
            element={
              localStorage.getItem("user_id") ? (
                <Gallery />
              ) : (
                <Navigate to="/register" replace />
              )
            }
          />
          <Route path="/albums" element={<Albums />} />
          <Route path="/vault" element={<Vault />} />
        </Routes>
      </div>
    </div>
  );
}


export default App
