class Mcpvault < Formula
  desc "CLI tool for managing MCP server configurations"
  homepage "https://github.com/mcpvault/mcpvault"
  url "https://github.com/mcpvault/mcpvault/releases/download/v0.1.0/mcpvault_0.1.0_darwin_amd64.tar.gz"
  sha256 "PLACEHOLDER_FOR_CHECKSUM"
  license "MIT"
  
  depends_on "go" => :build
  
  def install
    bin.install "mcpv"
  end
  
  test do
    system "#{bin}/mcpv", "--help"
  end
end 