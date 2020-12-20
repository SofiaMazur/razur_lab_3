const forumsClient = require('./forumsClient');

const testData = [
  ['addSite', 'Adding a new site (no args)', []],
  ['addSite', 'Adding a new site (no name)', [undefined, 'golang']],
  ['addSite', 'Adding a new site (no topic)', ['Gophers', undefined]],
  ['addSite', 'Adding a new site (empty name)', ['', 'Books']],
  ['addSite', 'Adding a new site (empty topic)', ['Gophers', '']],
  ['addSite', 'Adding a new site (name exists)', ['Book enjoyer', 'reading']],
  ['addSite', 'Adding a new site (topic exists)', ['Film Lover', 'Movies']],
  ['addSite', 'Adding a new site', ['Gophers', 'golang']],
  ['addUser', 'Adding a new user (no args)', []],
  ['addUser', 'Adding a new user (no name)', [undefined, ['golang']]],
  ['addUser', 'Adding a new user (no interests)', ['Barbara', undefined]],
  ['addUser', 'Adding a new user (empty string interests)', ['Barbara', ['Movies', '']]],
];

const sendTestResponses = async () => {
  const separator = '\n=========================================================\n';
  for (const [ method, comment, args ] of testData)
    try {
      console.log(comment, '\n');
      const responseFn = forumsClient[method];
      const res = await responseFn(...args);
      console.log('Result:')
      console.dir(res, { depth: null })
      console.log(separator)
    } catch (e) {
      console.error('Error:', e, '\n', separator);
    }
};

sendTestResponses();
