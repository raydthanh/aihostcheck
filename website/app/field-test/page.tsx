import type { Metadata } from "next";

const github = "https://github.com/raydthanh/aihostcheck";

export const metadata: Metadata = {
  title: "Community field test",
  description:
    "Run AIHostCheck on a real Windows, macOS, or Linux host and help improve evidence-based diagnostics for AI coding workflows.",
  alternates: {
    canonical: "https://aihostcheck.bond/field-test",
  },
  openGraph: {
    title: "Test AIHostCheck on a real developer machine",
    description:
      "A public, privacy-bounded field test for cross-OS host evidence—no signup, telemetry, or automatic report upload.",
    url: "https://aihostcheck.bond/field-test",
  },
};

function ArrowIcon() {
  return (
    <svg viewBox="0 0 18 18" aria-hidden="true">
      <path d="M4 14 14 4M6 4h8v8" />
    </svg>
  );
}

function CheckIcon() {
  return (
    <svg viewBox="0 0 20 20" aria-hidden="true">
      <path d="m4 10 4 4 8-9" />
    </svg>
  );
}

const environments = [
  "Windows · PowerShell, CMD, or Git Bash",
  "macOS · Intel or Apple silicon",
  "Linux · different distributions and shells",
  "Multiple Python or Node.js installations",
  "Docker, Podman, NVIDIA, CUDA, or local AI",
  "Restricted, virtualized, or containerized hosts",
];

const reviewPoints = [
  "OS, architecture, CPU, memory, and storage evidence",
  "The shell and the executable that wins on PATH",
  "Python, Node.js, Go, Java, Git, and package managers",
  "Container, GPU, NVIDIA driver, and CUDA visibility",
  "Unknown, unsupported, denied, and failed states",
  "Installation, checksums, and documentation friction",
];

export default function FieldTestPage() {
  return (
    <>
      <header className="site-header field-header">
        <a href="/" className="brand" aria-label="AIHostCheck home">
          <span className="brand-mark" aria-hidden="true">
            <span>&gt;_</span>
          </span>
          <span>AIHostCheck</span>
        </a>
        <nav aria-label="Field test navigation">
          <a href="#protocol">Protocol</a>
          <a href="#review">Review</a>
          <a href="#privacy">Privacy</a>
        </nav>
        <a className="header-link" href={github} target="_blank" rel="noreferrer">
          GitHub
          <ArrowIcon />
        </a>
      </header>

      <main>
        <section className="field-hero shell">
          <div className="field-kicker">
            <span className="pulse" />
            PUBLIC FIELD TEST · v0.1.0 · NO SIGNUP
          </div>
          <div className="field-hero-grid">
            <div>
              <h1>
                Help test what AI can safely know about a <em>host.</em>
              </h1>
            </div>
            <div className="field-hero-copy">
              <p>
                AIHostCheck is ready to leave hosted CI and meet real developer
                machines. Run it locally, inspect every result, and tell us
                where the evidence is wrong, ambiguous, or incomplete.
              </p>
              <div className="hero-actions">
                <a
                  className="button button-primary"
                  href={`${github}/releases/tag/v0.1.0`}
                  target="_blank"
                  rel="noreferrer"
                >
                  Download v0.1.0
                  <ArrowIcon />
                </a>
                <a
                  className="button button-secondary"
                  href={`${github}/blob/main/docs/FIELD_TESTING.md`}
                  target="_blank"
                  rel="noreferrer"
                >
                  Read the test guide
                </a>
              </div>
            </div>
          </div>
        </section>

        <section className="field-proof" aria-label="Field test boundaries">
          <div>
            <small>COLLECTION</small>
            <strong>Read-only</strong>
            <span>No network probes</span>
          </div>
          <div>
            <small>REPORTING</small>
            <strong>You review first</strong>
            <span>No automatic upload</span>
          </div>
          <div>
            <small>WHAT WE SEEK</small>
            <strong>Evidence</strong>
            <span>Not stars or endorsements</span>
          </div>
        </section>

        <section className="field-thesis shell">
          <div className="section-index">WHY REAL MACHINES MATTER</div>
          <div className="field-thesis-grid">
            <h2>Three operating systems are not three environments.</h2>
            <div>
              <p>
                Native CI proves that the project compiles and its automated
                tests pass on Windows, macOS, and Linux. It cannot reproduce the
                diversity of shells, permissions, package managers, runtimes,
                GPUs, or corporate restrictions found on real hosts.
              </p>
              <p>
                A plain laptop is useful. A complicated workstation is useful.
                The question is the same: can an AI consumer distinguish what
                is present from what could not be established?
              </p>
            </div>
          </div>
          <div className="environment-grid">
            {environments.map((environment) => (
              <div key={environment}>
                <CheckIcon />
                <span>{environment}</span>
              </div>
            ))}
          </div>
        </section>

        <section className="protocol-section" id="protocol">
          <div className="shell">
            <div className="section-index light">A SMALL, REPEATABLE PROTOCOL</div>
            <div className="protocol-heading">
              <h2>About ten minutes. No account. No report submission required.</h2>
              <p>
                Keep the report on your machine. Share only a minimal,
                reviewed excerpt when there is something the project can fix.
              </p>
            </div>
            <div className="protocol-grid">
              <article>
                <span>01</span>
                <h3>Download &amp; verify</h3>
                <p>
                  Choose the native archive for your platform and verify its
                  published SHA-256 checksum.
                </p>
              </article>
              <article>
                <span>02</span>
                <h3>Run both views</h3>
                <p>
                  Read the terminal report, then generate the versioned JSON
                  contract with <code>aihostcheck --json</code>.
                </p>
              </article>
              <article>
                <span>03</span>
                <h3>Compare locally</h3>
                <p>
                  Check the result against what you already know about the
                  machine and its active developer toolchain.
                </p>
              </article>
              <article>
                <span>04</span>
                <h3>Report the smallest gap</h3>
                <p>
                  If something is wrong or unclear, submit only the minimum
                  redacted evidence needed to reproduce it.
                </p>
              </article>
            </div>
          </div>
        </section>

        <section className="field-review shell" id="review">
          <div>
            <div className="section-index">WHAT TO REVIEW</div>
            <h2>Look for the point where an AI could make the wrong choice.</h2>
            <p>
              A missing version, ambiguous shell, or false absence can change
              the command an assistant proposes. Those details are the field
              test—not whether the page looks convincing.
            </p>
          </div>
          <ul>
            {reviewPoints.map((point) => (
              <li key={point}>
                <CheckIcon />
                {point}
              </li>
            ))}
          </ul>
        </section>

        <section className="privacy-callout" id="privacy">
          <div className="shell privacy-grid">
            <div>
              <div className="section-index light">PRIVACY BOUNDARY</div>
              <h2>Your report stays yours.</h2>
            </div>
            <div>
              <p>
                AIHostCheck has no telemetry and never uploads a report. Do not
                paste an unreviewed full report into an issue. Remove usernames,
                hostnames, private paths, IP addresses, organization details,
                and unrelated identifiers from any excerpt.
              </p>
              <a
                className="text-link dark-link"
                href={`${github}/blob/main/PRIVACY.md`}
                target="_blank"
                rel="noreferrer"
              >
                Inspect the data boundary
                <ArrowIcon />
              </a>
            </div>
          </div>
        </section>

        <section className="field-outcome shell">
          <div className="outcome-label">WHAT A GOOD FIELD TEST PRODUCES</div>
          <h2>Confirmed evidence, reproducible gaps, and better tests.</h2>
          <p>
            There is no target number of stars. One precise report that exposes
            a limitation is more valuable at this stage than an unqualified
            endorsement.
          </p>
          <div className="hero-actions centered">
            <a
              className="button button-primary"
              href={`${github}/issues/new?template=bug_report.yml`}
              target="_blank"
              rel="noreferrer"
            >
              Share a reviewed finding
              <ArrowIcon />
            </a>
            <a className="button button-secondary" href="/roadmap">
              See what evidence changes next
            </a>
          </div>
        </section>
      </main>

      <footer>
        <a href="/" className="brand footer-brand">
          <span className="brand-mark" aria-hidden="true">
            <span>&gt;_</span>
          </span>
          <span>AIHostCheck</span>
        </a>
        <p>Public field testing for trustworthy host evidence.</p>
        <div>
          <a href="/">Home</a>
          <a href="/roadmap">Roadmap</a>
          <a href={github} target="_blank" rel="noreferrer">
            GitHub
          </a>
        </div>
      </footer>
    </>
  );
}
