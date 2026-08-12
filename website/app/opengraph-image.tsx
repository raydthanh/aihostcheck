import { ImageResponse } from "next/og";

export const alt = "AIHostCheck — Host context AI should not guess";
export const size = { width: 1200, height: 630 };
export const contentType = "image/png";

export default function Image() {
  return new ImageResponse(
    (
      <div
        style={{
          width: "100%",
          height: "100%",
          display: "flex",
          flexDirection: "column",
          justifyContent: "space-between",
          background: "#07110f",
          color: "#ecf6f1",
          padding: "72px",
          fontFamily: "sans-serif",
        }}
      >
        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: "22px",
            fontSize: "32px",
            letterSpacing: "-1px",
          }}
        >
          <div
            style={{
              width: "54px",
              height: "54px",
              border: "4px solid #c7ff4a",
              display: "flex",
            }}
          />
          AIHostCheck
        </div>
        <div style={{ display: "flex", flexDirection: "column", gap: "20px" }}>
          <div
            style={{
              maxWidth: "980px",
              fontSize: "78px",
              lineHeight: 1.02,
              letterSpacing: "-4px",
              fontWeight: 700,
            }}
          >
            Host context AI should not guess.
          </div>
          <div style={{ color: "#a8b8b1", fontSize: "27px" }}>
            Open-source cross-OS diagnostics for AI development workflows.
          </div>
        </div>
        <div
          style={{
            display: "flex",
            color: "#c7ff4a",
            fontSize: "22px",
            fontFamily: "monospace",
          }}
        >
          Windows / macOS / Linux · Offline · Machine-readable
        </div>
      </div>
    ),
    size,
  );
}
