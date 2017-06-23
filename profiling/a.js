const http = require('http');

const hostname = '127.0.0.1';
const port = 8081;

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'text/html');
  var lines = ['<ul>']
  for (var i=0; i<1000;i++) {
    lines.push(`<li>${i+1}</li>`)
  }
  lines.push('</ul>')
  res.end(lines.join(''));
});

server.listen(port, hostname, () => {
  console.log(`Server running at http://${hostname}:${port}/`);
});