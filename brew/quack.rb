class Quack < Formula
  desc "Keep your IP address updated on Duck DNS"
  homepage "https://github.com/fcole90/quack"
  version "0.0.1"
  license "MIT"
  head "https://github.com/fcole90/quack.git", branch: "main"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args
  end

  test do
    system "#{bin}/quack", "set", "-token=\"my-token\"", "-domain=\"my-domain\"", "-timeinterval=300"
  end
end
