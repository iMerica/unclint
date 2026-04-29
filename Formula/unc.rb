class Unc < Formula
  desc "CLI linter for unc, corporate, stale, and tryhard language"
  homepage "https://github.com/iMerica/unclint"
  version "0.1.5"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.5/unc_v0.1.5_darwin_arm64.tar.gz"
      sha256 "6f3fd9291177de8b7ef1e69c1ada24cc8350da36086aca4f0fb58d276b5cbe6b"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.5/unc_v0.1.5_darwin_amd64.tar.gz"
      sha256 "cce54d88ad9d7c4b332d96f78b8652567bde709af3b6fa6661b1f25a0442c758"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.5/unc_v0.1.5_linux_arm64.tar.gz"
      sha256 "b54536bd2fe3f238d92f1eb04f8e707ad6aa67d392d878e9182340c3e149c046"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.5/unc_v0.1.5_linux_amd64.tar.gz"
      sha256 "2022478183ccb51a77bd1742cec894ac884ddf29b8703af53bfdc8d02cdcc972"
    end
  end

  def install
    bin.install "unc"
  end

  test do
    assert_match "unc", shell_output("#{bin}/unc --version")
  end
end
