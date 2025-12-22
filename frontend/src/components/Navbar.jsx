import { NavLink } from "react-router-dom";
import "./Navbar.css";

function Navbar() {
  return (
    <nav className="Navbar">
      <div className="NavLinks">
        <NavLink to="/">Gallery</NavLink>
        <NavLink to="/albums">Albums</NavLink>
        <NavLink to="/vault">Vault</NavLink>
      </div>
    </nav>
  );
}

export default Navbar;