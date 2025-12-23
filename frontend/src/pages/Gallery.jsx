import { useEffect, useState } from "react";
import { apiFetch } from "../api";

// Loads the user's media list from the backend when the Gallery page mounts
function Gallery() {
  const [items, setItems] = useState([]);
  const [error, setError] = useState("");
  const [thumbUrls, setThumbUrls] = useState({});

  useEffect(() => {
    async function loadMedia() {
      setError("");

      try {
        const res = await apiFetch("/media");
        const data = await res.json();

        const list = data.items || [];
        setItems(list);

        const urls = {};

        for (const it of list) {
          const fileRes = await apiFetch(`/media/${it.media_id}/file`);
          const blob = await fileRes.blob();
          const url = URL.createObjectURL(blob);
          urls[it.media_id] = url;
        }

        setThumbUrls(urls);
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

      <div style={{ display: "flex", justifyContent: "center" }}>
        <div
          style={{
            display: "grid",
            gridTemplateColumns: "repeat(3, 120px)",
            gap: 12,
          }}
        >
          {items.map((it) => {
            const url = thumbUrls[it.media_id];

            if (!url) {
              return (
                <div
                  key={it.media_id}
                  style={{
                    border: "2px solid black",
                    height: 120,
                    width: 120,
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  loading...
                </div>
              );
            }

            if (it.media_type === "video") {
              return (
                <video
                  key={it.media_id}
                  src={url}
                  controls
                  style={{
                    width: 120,
                    height: 120,
                    objectFit: "cover",
                    border: "2px solid black",
                  }}
                />
              );
            }

            return (
              <img
                key={it.media_id}
                src={url}
                alt=""
                style={{
                  width: 120,
                  height: 120,
                  objectFit: "cover",
                  border: "2px solid black",
                }}
              />
            );
          })}
        </div>
      </div>
    </div>
  );
}

export default Gallery;