class Unc < Formula
  desc "CLI linter for unc, corporate, stale, and tryhard language"
  homepage "https://github.com/iMerica/unclint"
  version "0.1.4"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.4/unc_v0.1.4_darwin_arm64.tar.gz"
      sha256 "caabe9612b3af03764006b56b0d723544865ecc348c7c60888d5d6dcb4c152e7"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.4/unc_v0.1.4_darwin_amd64.tar.gz"
      sha256 "be653c80407f3e1a0d0eb6162df26019ce1b828e69854388faff0ca056e3c2cd"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.4/unc_v0.1.4_linux_arm64.tar.gz"
      sha256 "b5e3297ae1813d56b2f812fed262e9062cb5a762b7a05330c94ccb59e2a993ea"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.4/unc_v0.1.4_linux_amd64.tar.gz"
      sha256 "449d51c6595101a595d04a897f00a887635e641da149646c2d716f51d3363a5a"
    end
  end

  def install
    bin.install "unc"
  end

  test do
    assert_match "unc", shell_output("#{bin}/unc --version")
  end
end
