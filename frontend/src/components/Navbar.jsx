import { NavLink } from "react-router-dom";

function Navbar() {
  return (
    <nav style={{ display: "flex", gap: "16px", padding: "12px", borderBottom: "1px solid #ddd" }}>
      <NavLink to="/">Gallery</NavLink>
      <NavLink to="/albums">Albums</NavLink>
      <NavLink to="/vault">Vault</NavLink>
    </nav>
  );
}

export default Navbar;