import { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { apiFetch } from "../api";

function Login() {
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
      const res = await apiFetch("/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      const data = await res.json();

      if (!data.user_id) {
        throw new Error("login succeeded but no user_id returned");
      }

      localStorage.setItem("user_id", String(data.user_id));
      navigate("/gallery", { replace: true });
    } catch (e2) {
      setError(e2.message || "login failed");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div
    style={{
        minHeight: "100vh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        marginTop: "-100px",
    }}
    >
    <div style={{ width: 360 }}>
      <h1 style={{ textAlign: "center" }}>Login</h1>

      {error ? <div style={{ marginBottom: 12 }}>{error}</div> : null}

      <form onSubmit={onSubmit} style={{ display: "flex", flexDirection: "column", gap: 12 }}>
        <label style={{ display: "flex", flexDirection: "column", gap: 6 }}>
          Username
          <input
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="username"
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
          {loading ? "Logging in..." : "Log in"}
        </button>
      </form>

      <div style={{ marginTop: 12, textAlign: "center" }}>
        Need an account? <Link to="/register">Register</Link>
      </div>
    </div>
    </div>
  );
}

export default Login;