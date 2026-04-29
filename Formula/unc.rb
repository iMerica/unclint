class Unc < Formula
  desc "CLI linter for unc, corporate, stale, and tryhard language"
  homepage "https://github.com/iMerica/unclint"
  version "0.1.8"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.8/unc_v0.1.8_darwin_arm64.tar.gz"
      sha256 "33ded0f37440e9b4d4de68adf05aed0deef56f0de3c6ea3b88293936b70a67d0"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.8/unc_v0.1.8_darwin_amd64.tar.gz"
      sha256 "be0294185a0960cb4be9a948b45a7cbf4829eb0954642a4463737de1c3f16f8c"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.8/unc_v0.1.8_linux_arm64.tar.gz"
      sha256 "a1fcce7260a5120147a3a99d7adf8e11e402c2578a235670714a203e416628f2"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.8/unc_v0.1.8_linux_amd64.tar.gz"
      sha256 "99f1275cdce5a2cdd805bae45821028737ab9b82f75052e029e174050938c2ca"
    end
  end

  def install
    bin.install "unc"
  end

  test do
    assert_match "unc", shell_output("#{bin}/unc --version")
  end
end
