import { Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Gallery from "./pages/Gallery";
import Albums from "./pages/Albums";
import Vault from "./pages/Vault";

function App() {
  return (
    <div>
      <Navbar />
      <div style={{ padding: "12px" }}>
        <Routes>
          <Route path="/" element={<Gallery />} />
          <Route path="/albums" element={<Albums />} />
          <Route path="/vault" element={<Vault />} />
        </Routes>
      </div>
    </div>
  );
}


export default App
