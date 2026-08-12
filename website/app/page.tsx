import { getProjectSnapshot } from "@/lib/github";

const github = "https://github.com/raydthanh/aihostcheck";

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

function formatDate(date: string) {
  return new Intl.DateTimeFormat("en", {
    year: "numeric",
    month: "short",
    day: "numeric",
  }).format(new Date(date));
}

export default async function Home() {
  const project = await getProjectSnapshot();

  const structuredData = {
    "@context": "https://schema.org",
    "@type": "SoftwareApplication",
    name: "AIHostCheck",
    applicationCategory: "DeveloperApplication",
    operatingSystem: "Windows, macOS, Linux",
    softwareVersion: project.release,
    url: "https://aihostcheck.bond",
    downloadUrl: project.releaseUrl,
    codeRepository: project.repoUrl,
    license: "https://www.apache.org/licenses/LICENSE-2.0",
    description:
      "Open-source, privacy-aware host diagnostics for AI development workflows.",
  };

  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(structuredData) }}
      />

      <header className="site-header">
        <a href="#top" className="brand" aria-label="AIHostCheck home">
          <span className="brand-mark" aria-hidden="true">
            <span>&gt;_</span>
          </span>
          <span>AIHostCheck</span>
        </a>
        <nav aria-label="Main navigation">
          <a href="#why">Why</a>
          <a href="#proof">Project</a>
          <a href="/field-test">Field test</a>
          <a href="/roadmap">Roadmap</a>
        </nav>
        <a className="header-link" href={github} target="_blank" rel="noreferrer">
          GitHub
          <ArrowIcon />
        </a>
      </header>

      <main id="top">
        <section className="hero shell">
          <div className="hero-copy">
            <div className="eyebrow">
              <span className="pulse" />
              Open source · early stage · shipping
            </div>
            <h1>
              Host context
              <br />
              AI should <em>not guess.</em>
            </h1>
            <p className="hero-lead">
              AIHostCheck gives GPT and AI coding agents evidence about a
              developer machine before they propose commands—across Windows,
              macOS, and Linux.
            </p>
            <div className="hero-actions">
              <a
                className="button button-primary"
                href={project.releaseUrl}
                target="_blank"
                rel="noreferrer"
              >
                Download {project.release}
                <ArrowIcon />
              </a>
              <a
                className="button button-secondary"
                href={github}
                target="_blank"
                rel="noreferrer"
              >
                Explore the source
              </a>
            </div>
            <div className="trust-line" aria-label="Core trust properties">
              <span>
                <CheckIcon />
                Read-only
              </span>
              <span>
                <CheckIcon />
                No telemetry
              </span>
              <span>
                <CheckIcon />
                Apache-2.0
              </span>
            </div>
          </div>

          <div className="terminal-wrap" aria-label="Example AIHostCheck report">
            <div className="terminal-glow" />
            <div className="terminal">
              <div className="terminal-bar">
                <div className="window-dots" aria-hidden="true">
                  <span />
                  <span />
                  <span />
                </div>
                <span>aihostcheck-report.json</span>
                <span className="terminal-state">verified</span>
              </div>
              <pre>
                <code>
                  <span className="line faint">{"{"}</span>
                  <span className="line">
                    {"  "}
                    <i>&quot;schema_version&quot;</i>:{" "}
                    <b>&quot;1.0.0&quot;</b>,
                  </span>
                  <span className="line">
                    {"  "}
                    <i>&quot;platform&quot;</i>:{" "}
                    <b>&quot;linux/amd64&quot;</b>,
                  </span>
                  <span className="line">
                    {"  "}
                    <i>&quot;capabilities&quot;</i>: {"{"}
                  </span>
                  <span className="line">
                    {"    "}
                    <i>&quot;python&quot;</i>: {"{"}
                  </span>
                  <span className="line">
                    {"      "}
                    <i>&quot;status&quot;</i>:{" "}
                    <strong>&quot;detected&quot;</strong>,
                  </span>
                  <span className="line">
                    {"      "}
                    <i>&quot;value&quot;</i>:{" "}
                    <b>&quot;Python 3.12.4&quot;</b>,
                  </span>
                  <span className="line">
                    {"      "}
                    <i>&quot;evidence&quot;</i>: [{"{"}
                  </span>
                  <span className="line">
                    {"        "}
                    <i>&quot;source&quot;</i>:{" "}
                    <b>&quot;executable_probe&quot;</b>
                  </span>
                  <span className="line faint">{"      }]"}</span>
                  <span className="line faint">{"    }"}</span>
                  <span className="line faint">{"  }"}</span>
                  <span className="line faint">{"}"}</span>
                </code>
              </pre>
              <div className="terminal-footer">
                <span>AI-readable</span>
                <span>Evidence-backed</span>
                <span>Offline</span>
              </div>
            </div>
          </div>
        </section>

        <section className="proof-strip" id="proof" aria-label="Project proof">
          <div className="proof-item">
            <small>Latest release</small>
            <strong>{project.release}</strong>
            <span>{formatDate(project.publishedAt)}</span>
          </div>
          <div className="proof-item">
            <small>Native targets</small>
            <strong>{project.packages || 6}</strong>
            <span>AMD64 + ARM64</span>
          </div>
          <div className="proof-item">
            <small>CI coverage</small>
            <strong>3 OS</strong>
            <span>Windows · macOS · Linux</span>
          </div>
          <div className="proof-item">
            <small>Data source</small>
            <strong>GitHub</strong>
            <span>
              {project.source === "github"
                ? "Live public project data"
                : "Last verified project data"}
            </span>
          </div>
        </section>

        <section className="problem shell" id="why">
          <div className="section-index">01 / THE GAP</div>
          <div className="section-heading">
            <h2>
              One wrong assumption.
              <br />
              Ten wrong commands.
            </h2>
            <p>
              A chat model cannot see the machine in front of you. Even an
              agent with terminal access can confuse absence with missing
              permissions, unsupported checks, or the wrong runtime on PATH.
            </p>
          </div>

          <div className="comparison">
            <article className="comparison-card before">
              <div className="card-label">WITHOUT HOST CONTEXT</div>
              <div className="chat-line">
                <span>AI</span>
                <p>Run <code>sudo apt install...</code></p>
              </div>
              <div className="chat-line user-line">
                <span>YOU</span>
                <p>I&apos;m on Windows.</p>
              </div>
              <div className="failure">Assumption → friction → lost time</div>
            </article>

            <article className="comparison-card after">
              <div className="card-label">WITH AIHOSTCHECK</div>
              <div className="context-row">
                <span>OS</span>
                <strong>Windows 11 / amd64</strong>
                <b>detected</b>
              </div>
              <div className="context-row">
                <span>Shell</span>
                <strong>PowerShell</strong>
                <b>detected</b>
              </div>
              <div className="context-row">
                <span>Package manager</span>
                <strong>winget</strong>
                <b>detected</b>
              </div>
              <div className="success">Evidence → relevant instruction</div>
            </article>
          </div>
        </section>

        <section className="how shell">
          <div className="section-index">02 / HOW IT WORKS</div>
          <div className="how-grid">
            <article>
              <span className="step-number">01</span>
              <h3>Run locally</h3>
              <p>
                One standalone binary performs bounded, read-only checks. No
                account, runtime, administrator privilege, or network probe.
              </p>
              <code>aihostcheck --json</code>
            </article>
            <article>
              <span className="step-number">02</span>
              <h3>Review the report</h3>
              <p>
                Inspect the JSON before sharing. Sensitive categories are
                excluded by design and reports are never uploaded.
              </p>
              <code>schema_version: 1.0.0</code>
            </article>
            <article>
              <span className="step-number">03</span>
              <h3>Give AI evidence</h3>
              <p>
                The agent sees explicit status, value, and evidence—and knows
                when to ask for a safe check instead of guessing.
              </p>
              <code>unknown ≠ absent</code>
            </article>
          </div>
        </section>

        <section className="contract shell">
          <div className="contract-copy">
            <div className="section-index">03 / WHY IT IS DIFFERENT</div>
            <h2>Not another system-info screenshot.</h2>
            <p>
              AIHostCheck is designed as a compatibility layer between a real
              host and an AI workflow. Its output is versioned, machine-readable,
              privacy-bounded, and explicit about uncertainty.
            </p>
            <a
              className="text-link"
              href={`${github}/blob/main/schema/report.schema.json`}
              target="_blank"
              rel="noreferrer"
            >
              Inspect the JSON Schema
              <ArrowIcon />
            </a>
          </div>
          <div className="status-list">
            {[
              ["detected", "Positive evidence exists"],
              ["not_detected", "Supported check found nothing"],
              ["unknown", "Evidence cannot establish the fact"],
              ["unsupported", "No supported check on this platform"],
              ["error", "A supported check failed"],
              ["permission_denied", "Access was blocked"],
            ].map(([status, meaning]) => (
              <div className="status-row" key={status}>
                <code>{status}</code>
                <span>{meaning}</span>
              </div>
            ))}
          </div>
        </section>

        <section className="roadmap shell" id="roadmap">
          <div className="section-index">04 / PROJECT TRAJECTORY</div>
          <div className="roadmap-title">
            <h2>Built honestly. Designed to grow.</h2>
            <p>
              AIHostCheck is an early-stage open-source project. The foundation
              ships today; broader validation and integrations are the next
              work—not claims we make in advance.
            </p>
          </div>

          <div className="roadmap-board">
            <article className="roadmap-column current">
              <div className="roadmap-label">
                <span />
                AVAILABLE NOW
              </div>
              <ul>
                <li>
                  <CheckIcon />
                  Native Windows, macOS, and Linux CLI
                </li>
                <li>
                  <CheckIcon />
                  Versioned JSON report contract
                </li>
                <li>
                  <CheckIcon />
                  Hardware, runtime, container, and GPU evidence
                </li>
                <li>
                  <CheckIcon />
                  Cross-OS CI and verified release checksums
                </li>
                <li>
                  <CheckIcon />
                  Privacy, security, and contribution policies
                </li>
              </ul>
            </article>
            <article className="roadmap-column next">
              <div className="roadmap-label">BUILDING NEXT</div>
              <ul>
                <li>
                  <span>01</span>
                  Real-device validation matrix
                </li>
                <li>
                  <span>02</span>
                  Signed and easier installation paths
                </li>
                <li>
                  <span>03</span>
                  Deeper AI and GPU runtime diagnostics
                </li>
                <li>
                  <span>04</span>
                  Reusable adapters for agent workflows
                </li>
                <li>
                  <span>05</span>
                  Local report viewer and redaction assistance
                </li>
              </ul>
            </article>
          </div>
          <a className="text-link roadmap-link" href="/roadmap">
            See milestones and completion evidence
            <ArrowIcon />
          </a>
        </section>

        <section className="partners shell">
          <div className="partner-card">
            <div>
              <div className="section-index light">05 / ECOSYSTEM SUPPORT</div>
              <h2>What infrastructure support unlocks.</h2>
            </div>
            <p>
              Developer program credits and tooling help an early open-source
              project turn a credible foundation into a tested public utility:
              more real-device coverage, safer releases, dependable hosting,
              and faster community feedback.
            </p>
            <div className="unlock-grid">
              <span>Cross-platform testing</span>
              <span>Release signing</span>
              <span>Documentation hosting</span>
              <span>Observability</span>
            </div>
          </div>
        </section>

        <section className="final-cta shell">
          <div className="cta-kicker">THE SOURCE OF TRUTH IS PUBLIC</div>
          <h2>
            Read the code.
            <br />
            Run the release.
            <br />
            Challenge the evidence.
          </h2>
          <div className="hero-actions centered">
            <a
              className="button button-primary"
              href={github}
              target="_blank"
              rel="noreferrer"
            >
              View on GitHub
              <ArrowIcon />
            </a>
            <a
              className="button button-secondary"
              href={`${github}/blob/main/docs/USING_WITH_AI.md`}
              target="_blank"
              rel="noreferrer"
            >
              Use with an AI agent
            </a>
          </div>
        </section>
      </main>

      <footer>
        <a href="#top" className="brand footer-brand">
          <span className="brand-mark" aria-hidden="true">
            <span>&gt;_</span>
          </span>
          <span>AIHostCheck</span>
        </a>
        <p>
          Open-source cross-OS diagnostics for AI developer environments.
        </p>
        <div>
          <a href={github} target="_blank" rel="noreferrer">
            GitHub
          </a>
          <a href="/field-test">Field test</a>
          <a href="/roadmap">Roadmap</a>
          <a
            href={`${github}/blob/main/PRIVACY.md`}
            target="_blank"
            rel="noreferrer"
          >
            Privacy
          </a>
          <a
            href={`${github}/blob/main/SECURITY.md`}
            target="_blank"
            rel="noreferrer"
          >
            Security
          </a>
        </div>
      </footer>
    </>
  );
}
