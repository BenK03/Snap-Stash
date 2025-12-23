import { Routes, Route, Navigate } from "react-router-dom";
import Gallery from "./pages/Gallery";
import Login from "./pages/Login";
import Register from "./pages/Register";

function App() {
  return (
      <div style={{ padding: "12px" }}>
        <Routes>
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />

          <Route
            path="/gallery"
            element={
              localStorage.getItem("user_id")
                ? <Gallery />
                : <Navigate to="/register" replace />
            }
          />

          {/* default route */}
          <Route
            path="/"
            element={<Navigate to="/register" replace />}
          />
        </Routes>
      </div>
  );
}

export default App;
