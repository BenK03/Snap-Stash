import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { apiFetch } from "../api";

function Register() {
  const navigate = useNavigate();

  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  async function onSubmit(e) {
    e.preventDefault();
    setError("");

    if (loading) {
      return;
    }

    setLoading(true);

    try {
      const res = await apiFetch("/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      if (!res.ok) {
        const msg = await res.text();
        throw new Error(msg || "register failed");
      }

      const data = await res.json();
      console.log("REGISTER RESPONSE:", data);

      const userId = data.user_id || data.userID || data.id;
      if (!userId) {
        throw new Error("register succeeded but no user_id returned");
      }

      localStorage.setItem("user_id", String(userId));
      navigate("/");
    } catch (e2) {
      setError(e2.message || "register failed");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div style={{ maxWidth: 360, margin: "0 auto" }}>
      <h1 style={{ textAlign: "center" }}>Register</h1>

      {error ? <div style={{ marginBottom: 12 }}>{error}</div> : null}

      <form onSubmit={onSubmit} style={{ display: "flex", flexDirection: "column", gap: 12 }}>
        <label style={{ display: "flex", flexDirection: "column", gap: 6 }}>
          Username
          <input
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="demo"
          />
        </label>

        <label style={{ display: "flex", flexDirection: "column", gap: 6 }}>
          Password
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="password"
          />
        </label>

        <button type="submit" disabled={loading}>
          {loading ? "Creating..." : "Create account"}
        </button>
      </form>

      <div style={{ marginTop: 12, textAlign: "center" }}>
        Already have an account? <Link to="/login">Login</Link>
      </div>
    </div>
  );
}

export default Register;