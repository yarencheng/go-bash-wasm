import fs from 'fs';
import path from 'path';
import { fileURLToPath } from 'url';

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const sitemapPath = path.resolve(__dirname, '../static/sitemap.xml');

if (!fs.existsSync(sitemapPath)) {
	console.error(`Sitemap not found at ${sitemapPath}`);
	process.exit(1);
}

const content = fs.readFileSync(sitemapPath, 'utf8');
const date = new Date().toISOString().split('T')[0];

const updated = content.replace(/<lastmod>.*<\/lastmod>/, `<lastmod>${date}</lastmod>`);

if (content !== updated) {
	fs.writeFileSync(sitemapPath, updated);
	console.log(`Successfully updated lastmod in sitemap.xml to ${date}`);
} else {
	console.log('Sitemap is already up to date.');
}
