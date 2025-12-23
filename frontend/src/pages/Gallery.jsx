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
      <div>Items loaded: {items.length}</div>
    </div>
  );
}

export default Gallery;