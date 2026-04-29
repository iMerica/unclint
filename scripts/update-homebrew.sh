#!/usr/bin/env bash
set -euo pipefail

: "${VERSION:?VERSION is required}"

OWNER_REPO="${GITHUB_REPOSITORY}"

DARWIN_ARM64_SHA="$(grep "darwin_arm64.tar.gz" dist/SHA256SUMS.txt | awk '{print $1}')"
DARWIN_AMD64_SHA="$(grep "darwin_amd64.tar.gz" dist/SHA256SUMS.txt | awk '{print $1}')"
LINUX_AMD64_SHA="$(grep "linux_amd64.tar.gz" dist/SHA256SUMS.txt | awk '{print $1}')"
LINUX_ARM64_SHA="$(grep "linux_arm64.tar.gz" dist/SHA256SUMS.txt | awk '{print $1}')"

mkdir -p Formula

cat > Formula/unc.rb <<EOF2
class Unc < Formula
  desc "CLI linter for unc, corporate, stale, and tryhard language"
  homepage "https://github.com/${OWNER_REPO}"
  version "${VERSION#v}"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/${OWNER_REPO}/releases/download/${VERSION}/unc_${VERSION}_darwin_arm64.tar.gz"
      sha256 "${DARWIN_ARM64_SHA}"
    else
      url "https://github.com/${OWNER_REPO}/releases/download/${VERSION}/unc_${VERSION}_darwin_amd64.tar.gz"
      sha256 "${DARWIN_AMD64_SHA}"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/${OWNER_REPO}/releases/download/${VERSION}/unc_${VERSION}_linux_arm64.tar.gz"
      sha256 "${LINUX_ARM64_SHA}"
    else
      url "https://github.com/${OWNER_REPO}/releases/download/${VERSION}/unc_${VERSION}_linux_amd64.tar.gz"
      sha256 "${LINUX_AMD64_SHA}"
    end
  end

  def install
    bin.install "unc"
  end

  test do
    assert_match "unc", shell_output("#{bin}/unc --version")
  end
end
EOF2

git config user.name "github-actions[bot]"
git config user.email "github-actions[bot]@users.noreply.github.com"
git add Formula/unc.rb
git commit -m "Update unc formula to ${VERSION} [skip ci]" || exit 0
git push origin HEAD:master