// Scry Info.  All rights reserved.
// license that can be found in the license file.

require('eventsource-polyfill');
let hotClient = require('webpack-hot-middleware/client?noInfo=true&reload=true');

hotClient.subscribe(function (event) {
  if (event.action === 'reload') {
    window.location.reload();
  }
});
