var path = require("path");

var root = path.join(__dirname);

var config = {
  rootDir: root,
  // Targets ========================================================
  serveDir: path.join(root, ".serve"),
  distDir: path.join(root, "dist"),
  clientManifestFile: "manifest.webpack.json",
  clientStatsFile: "stats.webpack.json",

  // Source Directory ===============================================
  srcDir: path.join(root, "app"),
  srcServerDir: path.join(root, "server"),

  // HTML Layout ====================================================
  srcHtmlLayout: path.join(root, "app", "index.html"),

  // Site Config ====================================================
  siteTitle: "Funda",
  siteDescription: "Dashboard",
  siteCannonicalUrl: "http://localhost:4100",
  siteKeywords: "Funda dashboard",
};

module.exports = config;
