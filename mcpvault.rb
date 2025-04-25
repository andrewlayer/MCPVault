class Mcpvault < Formula
  desc "CLI tool for managing MCP server configurations"
  homepage "https://github.com/andrewlayer/MCPVault"
  
  # For initial release - using the repository itself rather than a release archive
  url "https://github.com/andrewlayer/MCPVault.git", 
      tag:      "v0.1.0",
      revision: "HEAD"
  license "MIT"
  
  depends_on "go" => :build
  
  def install
    system "go", "build", *std_go_args(output: bin/"mcpv")
  end
  
  test do
    system "#{bin}/mcpv", "--version"
  end
end 