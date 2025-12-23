import { useEffect, useState } from "react";
import { apiFetch } from "../api";

// Loads the user's media list from the backend when the Gallery page mounts
function Gallery() {
  const [items, setItems] = useState([]);
  const [error, setError] = useState("");

  useEffect(() => {
    async function loadMedia() {
      setError("");

      try {
        const res = await apiFetch("/media");
        const data = await res.json();
        setItems(data.items || []);
      } catch (e) {
        setError(e.message || "failed to load media");
      }
    }

    loadMedia();
  }, []);

  return (
    <div>
      <h1>Gallery</h1>

      {error ? <div>{error}</div> : null}

      <div
        style={{
          display: "grid",
          gridTemplateColumns: "repeat(3, 1fr)",
          gap: 12,
        }}
      >
        {items.map((it) => {
          return (
            <div
              key={it.media_id}
              style={{
                border: "1px solid #ddd",
                height: 180,
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
                flexDirection: "column",
                gap: 6,
              }}
            >
              <div>id: {it.media_id}</div>
              <div>{it.media_type}</div>
            </div>
          );
        })}
      </div>
    </div>
  );
}

export default Gallery;