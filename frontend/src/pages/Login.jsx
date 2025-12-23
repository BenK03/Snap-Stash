import { useState } from "react";

function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  function onSubmit(e) {
    e.preventDefault();
    // TODO: call backend login
  }

  return (
    <div style={{ maxWidth: 360, margin: "0 auto" }}>
      <h1 style={{ textAlign: "center" }}>Login</h1>

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

        <button type="submit">Log in</button>
      </form>
    </div>
  );
}

export default Login;