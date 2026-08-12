import type { Metadata } from "next";

const github = "https://github.com/raydthanh/aihostcheck";

export const metadata: Metadata = {
  title: "Roadmap",
  description:
    "The evidence-based AIHostCheck roadmap: shipped foundation, measurable next outcomes, explicit non-goals, and how ecosystem support would be used.",
  alternates: {
    canonical: "/roadmap",
  },
  openGraph: {
    title: "AIHostCheck roadmap — evidence before promises",
    description:
      "A public, measurable roadmap for trustworthy host diagnostics across Windows, macOS, and Linux.",
    url: "/roadmap",
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

const phases = [
  {
    index: "00",
    status: "SHIPPED",
    title: "A trustworthy foundation",
    summary:
      "The first release proves the core contract: bounded local collection, explicit uncertainty, and native output an AI consumer can parse.",
    outcomes: [
      "Native Windows, macOS, and Linux CLI",
      "Versioned JSON contract and human-readable output",
      "Six release archives with SHA-256 checksums",
      "Cross-OS CI, privacy, security, and AI workflow guides",
    ],
    evidence: "v0.1.0 release",
    evidenceUrl: `${github}/releases/tag/v0.1.0`,
  },
  {
    index: "01",
    status: "NEXT",
    title: "Reliability on real machines",
    summary:
      "Move beyond hosted runners with a privacy-safe validation matrix and turn every confirmed mismatch into a regression test.",
    outcomes: [
      "Redacted fixture format for real-device reports",
      "Representative OS and architecture coverage",
      "Documented limitations for blocked or unavailable evidence",
      "Compatibility and provenance rules covered by tests",
    ],
    evidence: "Track execution issues",
    evidenceUrl: `${github}/issues`,
  },
  {
    index: "02",
    status: "FOLLOWING",
    title: "Safer distribution and AI workflows",
    summary:
      "Make reports easier to verify, inspect, redact, and reuse without requiring an account or uploading host data.",
    outcomes: [
      "Maintainable artifact-signing policy before signed releases",
      "Low-friction installation paths evaluated per OS family",
      "Local report review and redaction assistance",
      "Agent adapters with explicit unknown and failure behavior",
    ],
    evidence: "Read the normative roadmap",
    evidenceUrl: `${github}/blob/main/docs/ROADMAP.md`,
  },
  {
    index: "03",
    status: "EVIDENCE REQUIRED",
    title: "Specialized diagnostic packs",
    summary:
      "Deeper GPU, accelerator, local-AI runtime, editor, and support-bundle work enters the release plan only after real workflow evidence exists.",
    outcomes: [
      "NVIDIA driver and toolkit compatibility evidence",
      "Additional accelerator and local-runtime ecosystems",
      "Capability profiles for Python, containers, and GPU work",
      "Plugin model only if the core stays small and auditable",
    ],
    evidence: "Propose a bounded feature",
    evidenceUrl: `${github}/issues/new?template=feature_request.yml`,
  },
];

const supportRows = [
  [
    "Cross-platform compute",
    "Broader OS and architecture validation",
    "Fixtures, tests, and a public coverage matrix",
  ],
  [
    "Secure delivery",
    "Signing policy and verifiable native artifacts",
    "Documented process and signed releases",
  ],
  [
    "Hosting & observability",
    "Reliable docs and faster production diagnosis",
    "Public fixes and deployment history",
  ],
  [
    "AI developer tooling",
    "Validate adapters across agent workflows",
    "Open examples and compatibility notes",
  ],
];

export default function RoadmapPage() {
  return (
    <>
      <header className="site-header roadmap-header">
        <a href="/" className="brand" aria-label="AIHostCheck home">
          <span className="brand-mark" aria-hidden="true">
            <span>&gt;_</span>
          </span>
          <span>AIHostCheck</span>
        </a>
        <nav aria-label="Roadmap navigation">
          <a href="#phases">Milestones</a>
          <a href="#support">Support</a>
          <a href="#boundaries">Boundaries</a>
        </nav>
        <a className="header-link" href={github} target="_blank" rel="noreferrer">
          GitHub
          <ArrowIcon />
        </a>
      </header>

      <main>
        <section className="roadmap-hero shell">
          <div className="roadmap-kicker">
            <span className="pulse" />
            PUBLIC ROADMAP · UPDATED WITH PROJECT EVIDENCE
          </div>
          <h1>
            A roadmap
            <br />
            you can <em>audit.</em>
          </h1>
          <div className="roadmap-intro">
            <p>
              AIHostCheck does not count intentions as progress. Every milestone
              needs visible acceptance evidence in code, tests, documentation,
              an issue, or a release.
            </p>
            <a
              className="button button-secondary"
              href={`${github}/blob/main/docs/ROADMAP.md`}
              target="_blank"
              rel="noreferrer"
            >
              Read the source roadmap
              <ArrowIcon />
            </a>
          </div>
        </section>

        <section className="roadmap-proof" aria-label="Roadmap trust rules">
          <div>
            <small>CURRENT RELEASE</small>
            <strong>v0.1.0</strong>
            <span>Working early foundation</span>
          </div>
          <div>
            <small>VALIDATION STAGE</small>
            <strong>Early</strong>
            <span>Real-device coverage is next</span>
          </div>
          <div>
            <small>COMPLETION RULE</small>
            <strong>Public evidence</strong>
            <span>Code · tests · docs · release</span>
          </div>
        </section>

        <section className="roadmap-principles shell">
          <div className="section-index">HOW PRIORITIES MOVE</div>
          <div className="principle-grid">
            <article>
              <span>01</span>
              <h2>Incorrect diagnosis</h2>
              <p>
                A reproducible report reveals a wrong or ambiguous result that
                could make an AI choose the wrong instruction.
              </p>
            </article>
            <article>
              <span>02</span>
              <h2>Coverage gap</h2>
              <p>
                Real hardware or permissions expose evidence that hosted CI
                cannot validate safely.
              </p>
            </article>
            <article>
              <span>03</span>
              <h2>Bounded evidence</h2>
              <p>
                A contribution adds cross-OS value through a documented,
                read-only source with a clear privacy boundary.
              </p>
            </article>
          </div>
        </section>

        <section className="phase-section shell" id="phases">
          <div className="section-index">RELEASE TRAJECTORY</div>
          <div className="phase-heading">
            <h2>From foundation to reusable compatibility layer.</h2>
            <p>
              Sequence matters: first prove the evidence on real hosts, then
              improve distribution, then expand integrations.
            </p>
          </div>

          <div className="phase-list">
            {phases.map((phase) => (
              <article className="phase-card" key={phase.index}>
                <div className="phase-index">{phase.index}</div>
                <div className="phase-body">
                  <div className="phase-status">{phase.status}</div>
                  <h3>{phase.title}</h3>
                  <p>{phase.summary}</p>
                </div>
                <div className="phase-outcomes">
                  <small>COMPLETION EVIDENCE</small>
                  <ul>
                    {phase.outcomes.map((outcome) => (
                      <li key={outcome}>
                        <CheckIcon />
                        {outcome}
                      </li>
                    ))}
                  </ul>
                  <a href={phase.evidenceUrl} target="_blank" rel="noreferrer">
                    {phase.evidence}
                    <ArrowIcon />
                  </a>
                </div>
              </article>
            ))}
          </div>
        </section>

        <section className="support-section" id="support">
          <div className="shell">
            <div className="section-index light">SUPPORT → OUTPUT</div>
            <div className="support-heading">
              <h2>Resources must leave public evidence.</h2>
              <p>
                Infrastructure support accelerates validation and delivery. It
                is never presented as adoption, a customer relationship, or an
                endorsement.
              </p>
            </div>
            <div className="support-table" role="table" aria-label="How support is used">
              <div className="support-table-head" role="row">
                <span role="columnheader">Support area</span>
                <span role="columnheader">Project use</span>
                <span role="columnheader">Public evidence</span>
              </div>
              {supportRows.map(([area, use, evidence]) => (
                <div className="support-table-row" role="row" key={area}>
                  <strong role="cell">{area}</strong>
                  <span role="cell">{use}</span>
                  <span role="cell">{evidence}</span>
                </div>
              ))}
            </div>
          </div>
        </section>

        <section className="boundary-section shell" id="boundaries">
          <div>
            <div className="section-index">NON-NEGOTIABLE BOUNDARIES</div>
            <h2>Growth without changing the trust model.</h2>
          </div>
          <ul>
            <li>No remote command execution</li>
            <li>No telemetry or background monitoring</li>
            <li>No automatic report upload</li>
            <li>No account or hosted report database</li>
            <li>No silent software installation or reconfiguration</li>
            <li>No “absent” claim when evidence is unavailable</li>
          </ul>
        </section>

        <section className="roadmap-cta shell">
          <div className="cta-kicker">ROADMAPS EARN TRUST THROUGH EXECUTION</div>
          <h2>Inspect the plan. Challenge a gap. Follow the evidence.</h2>
          <div className="hero-actions centered">
            <a
              className="button button-primary"
              href={`${github}/issues`}
              target="_blank"
              rel="noreferrer"
            >
              View execution issues
              <ArrowIcon />
            </a>
            <a className="button button-secondary" href="/">
              Back to AIHostCheck
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
        <p>Evidence-based roadmap for an early-stage open-source project.</p>
        <div>
          <a href="/">Home</a>
          <a href={github} target="_blank" rel="noreferrer">
            GitHub
          </a>
          <a
            href={`${github}/blob/main/docs/ROADMAP.md`}
            target="_blank"
            rel="noreferrer"
          >
            Roadmap source
          </a>
        </div>
      </footer>
    </>
  );
}
