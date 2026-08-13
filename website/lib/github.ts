type GitHubRepository = {
  stargazers_count: number;
  updated_at: string;
  html_url: string;
};

type GitHubAsset = {
  name: string;
  download_count: number;
};

type GitHubRelease = {
  tag_name: string;
  published_at: string;
  html_url: string;
  assets: GitHubAsset[];
};

export type ProjectSnapshot = {
  release: string;
  publishedAt: string;
  repoUpdatedAt: string;
  stars: number;
  packages: number;
  repoUrl: string;
  releaseUrl: string;
  source: "github" | "fallback";
};

const fallback: ProjectSnapshot = {
  release: "v0.2.1",
  publishedAt: "2026-08-13T05:16:16Z",
  repoUpdatedAt: "2026-08-13T05:15:10Z",
  stars: 44,
  packages: 6,
  repoUrl: "https://github.com/raydthanh/aihostcheck",
  releaseUrl: "https://github.com/raydthanh/aihostcheck/releases/latest",
  source: "fallback",
};

export async function getProjectSnapshot(): Promise<ProjectSnapshot> {
  const headers = {
    Accept: "application/vnd.github+json",
    "X-GitHub-Api-Version": "2022-11-28",
  };

  try {
    const [repoResponse, releaseResponse] = await Promise.all([
      fetch("https://api.github.com/repos/raydthanh/aihostcheck", {
        headers,
        next: { revalidate: 900 },
      }),
      fetch("https://api.github.com/repos/raydthanh/aihostcheck/releases/latest", {
        headers,
        next: { revalidate: 900 },
      }),
    ]);

    if (!repoResponse.ok || !releaseResponse.ok) {
      return fallback;
    }

    const [repo, release] = (await Promise.all([
      repoResponse.json(),
      releaseResponse.json(),
    ])) as [GitHubRepository, GitHubRelease];

    return {
      release: release.tag_name,
      publishedAt: release.published_at,
      repoUpdatedAt: repo.updated_at,
      stars: repo.stargazers_count,
      packages: release.assets.filter((asset) =>
        /_(linux|darwin|windows)_(amd64|arm64)\.(tar\.gz|zip)$/.test(
          asset.name,
        ),
      ).length,
      repoUrl: repo.html_url,
      releaseUrl: release.html_url,
      source: "github",
    };
  } catch {
    return fallback;
  }
}
