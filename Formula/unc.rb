class Unc < Formula
  desc "CLI linter for unc, corporate, stale, and tryhard language"
  homepage "https://github.com/iMerica/unclint"
  version "0.1.7"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.7/unc_v0.1.7_darwin_arm64.tar.gz"
      sha256 "b2f1f1418ff4e178587262ce72b477ddff2c6bf637b05852fe5fe773aa60559f"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.7/unc_v0.1.7_darwin_amd64.tar.gz"
      sha256 "6f0c523ae312277f0e504ee6f162c68fa1fa8e1999caa724d46d10a452d14080"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.7/unc_v0.1.7_linux_arm64.tar.gz"
      sha256 "964c6a2da8900610080d97df688f94ec6fb725706cb26f9409ea68338fb25dca"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.7/unc_v0.1.7_linux_amd64.tar.gz"
      sha256 "8b3f724b2e1b907219a65ef95cbd25d4eb9d3f2b849c276631f2811354847e68"
    end
  end

  def install
    bin.install "unc"
  end

  test do
    assert_match "unc", shell_output("#{bin}/unc --version")
  end
end
