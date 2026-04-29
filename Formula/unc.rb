class Unc < Formula
  desc "CLI linter for unc, corporate, stale, and tryhard language"
  homepage "https://github.com/iMerica/unclint"
  version "0.1.6"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.6/unc_v0.1.6_darwin_arm64.tar.gz"
      sha256 "c011591ec5dbabb65dc69c0107a9bf066ee559b826bc9d4475b2b602031cf0a1"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.6/unc_v0.1.6_darwin_amd64.tar.gz"
      sha256 "fb6e78ddfd355b5310cabfdfcefe2cc2c231a0ec8d55a607e061fd4a8b0f7f71"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/iMerica/unclint/releases/download/v0.1.6/unc_v0.1.6_linux_arm64.tar.gz"
      sha256 "9a8d50ae501ca16123e1c2b7cb654ec562f8ebbc5bb97ee07ac1e180bae329d3"
    else
      url "https://github.com/iMerica/unclint/releases/download/v0.1.6/unc_v0.1.6_linux_amd64.tar.gz"
      sha256 "e42fe480b6401e18df153e897a33f11061fb68ed6e911d3824de6ae1a045cfbf"
    end
  end

  def install
    bin.install "unc"
  end

  test do
    assert_match "unc", shell_output("#{bin}/unc --version")
  end
end
