const forumsClient = require('./siteUser');

const testData = [
  ['addSite', 'Adding a new site (no topic)', []],
  ['addSite', 'Adding a new site (empty topic)', [undefined, 'js']],
  ['addSite', 'Adding a new site (no topic)', ['no', undefined]],
  ['addSite', 'Adding a new site (empty name)', ['', 'magazines']],
  ['addSite', 'Adding a new site (no name)', ['no', '']],
  ['addSite', 'Adding a new site (name exists)', ['Magazines', 'reading']],
  ['addSite', 'Adding a new site (topic exists)', ['TV-show', 'Films']],
  ['addSite', 'Adding a new site', ['Gophers', 'golang']],
  ['addUser', 'Adding a new user (no args)', []],
  ['addUser', 'Adding a new user (no name)', [undefined, ['golang']]],
  ['addUser', 'Adding a new user (no topic)', ['Dr. Strange', undefined]],
  ['addUser', 'Adding a new user (empty string interests)', ['Dr. Strange', ['Films', '']]],
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
