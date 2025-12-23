const API_BASE = "/api";

// Calls the backend API to handle media-related operations
export async function apiFetch(path, options = {}) {
  const res = await fetch(`${API_BASE}${path}`, {
    ...options,
    headers: {
      "X-User-ID": "1",
      ...(options.headers || {}),
    },
  });

  if (!res.ok) {
    throw new Error(`API error ${res.status}`);
  }

  return res;
}