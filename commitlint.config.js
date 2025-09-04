/** @type { import('cz-git').UserConfig } */
export default {
  prompt: {
    useEmoji: true,
    scopes: ['deps', 'dev'],
    allowEmptyScopes: true,
    allowCustomScopes: false,
    markBreakingChangeMode: true,
    allowBreakingChanges: ['feat', 'fix'],
    issuePrefixes: [
      { value: 'closed', name: 'closed:   ISSUES has been processed' },
    ],
  },
}
