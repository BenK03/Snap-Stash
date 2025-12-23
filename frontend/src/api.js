const API_BASE = "/api";

// apiFetch(path, options)
// - Sends requests to the backend under /api
// - Automatically adds the X-User-ID header if the user is logged in
export async function apiFetch(path, options = {}) {
  const userId = localStorage.getItem("user_id");

  const headers = {
    ...(options.headers || {}),
  };

  if (userId) {
    headers["X-User-ID"] = userId;
  }

  const res = await fetch(`${API_BASE}${path}`, {
    ...options,
    headers,
  });

  if (!res.ok) {
    const msg = await res.text();
    throw new Error(msg || `API error ${res.status}`);
  }

  return res;
}