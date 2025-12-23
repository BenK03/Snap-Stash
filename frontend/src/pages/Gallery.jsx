import { useEffect, useState } from "react";
import { apiFetch } from "../api";

// Loads the user's media list from the backend when the Gallery page mounts
function Gallery() {
  const [items, setItems] = useState([]);
  const [error, setError] = useState("");
  const [thumbUrls, setThumbUrls] = useState({});
  const [selected, setSelected] = useState(null);
  const [uploadFile, setUploadFile] = useState(null);
  const [uploading, setUploading] = useState(false);
  const [uploadError, setUploadError] = useState("");

  // Helper to load media and set state
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

  useEffect(() => {
    loadMedia();

    function onKeyDown(e) {
      if (e.key === "Escape") {
        setSelected(null);
      }
    }

    window.addEventListener("keydown", onKeyDown);

    return () => {
      window.removeEventListener("keydown", onKeyDown);
    };
  }, []);

  async function onUpload() {
    setUploadError("");

    if (!uploadFile) {
      setUploadError("pick a file first");
      return;
    }

    if (uploading) {
      return;
    }

    setUploading(true);

    try {
      const form = new FormData();
      form.append("file", uploadFile);

      await apiFetch("/media/upload", {
        method: "POST",
        body: form,
      });

      setUploadFile(null);

      // refresh gallery so the new item appears
      await loadMedia();
    } catch (e) {
      setUploadError(e.message || "upload failed");
    } finally {
      setUploading(false);
    }
  }

  return (
    <div>
      <h1 style={{ textAlign: "center" }}>Gallery</h1>

      <div style={{ display: "flex", justifyContent: "center", margin: "12px 0", gap: 12 }}>
        <input
          type="file"
          accept="image/*,video/*"
          onChange={(e) =>
            setUploadFile(e.target.files && e.target.files[0] ? e.target.files[0] : null)
          }
        />

        <button onClick={onUpload} disabled={uploading}>
          {uploading ? "Uploading..." : "Upload"}
        </button>
      </div>

      {uploadError ? (
        <div style={{ textAlign: "center", marginBottom: 12 }}>{uploadError}</div>
      ) : null}

      {error ? <div>{error}</div> : null}

      <div style={{ display: "flex", justifyContent: "center" }}>
        <div
          style={{
            display: "grid",
            gridTemplateColumns: "repeat(3, 120px)",
            gap: 70,
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
                  muted
                  playsInline
                  preload="metadata"
                  onClick={() => setSelected({
                    media_id: it.media_id,
                    media_type: it.media_type,
                    url,
                  })}
                  style={{
                    width: 120,
                    height: 120,
                    objectFit: "cover",
                    border: "2px solid black",
                    cursor: "pointer",
                  }}
                />
              );
            }

            return (
              <img
                key={it.media_id}
                src={url}
                alt=""
                onClick={() => setSelected({
                  media_id: it.media_id,
                  media_type: it.media_type,
                  url,
                })}
                style={{
                  width: 120,
                  height: 120,
                  objectFit: "cover",
                  border: "2px solid black",
                  cursor: "pointer",
                }}
              />
            );
          })}
        </div>
      </div>
      {selected ? (
        <div
          onClick={() => setSelected(null)}
          style={{
            position: "fixed",
            inset: 0,
            background: "rgba(0,0,0,0.85)",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            padding: 24,
            zIndex: 9999,
          }}
        >
          <div
            onClick={(e) => e.stopPropagation()}
            style={{
              maxWidth: "90vw",
              maxHeight: "90vh",
            }}
          >
            {selected.media_type === "video" ? (
              <video
                src={selected.url}
                controls
                autoPlay
                style={{
                  maxWidth: "90vw",
                  maxHeight: "90vh",
                }}
              />
            ) : (
              <img
                src={selected.url}
                alt=""
                style={{
                  maxWidth: "90vw",
                  maxHeight: "90vh",
                  objectFit: "contain",
                  display: "block",
                }}
              />
            )}
          </div>
        </div>
      ) : null}
    </div>
  );
}

export default Gallery;