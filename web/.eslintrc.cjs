/**
 *  @type {import('eslint').ESLint.ConfigData}
 */
module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  plugins: ['perfectionist', 'unused-imports', '@typescript-eslint', 'prettier'],
  extends: ['airbnb', 'airbnb-typescript', 'airbnb/hooks', 'prettier'],
  parserOptions: {
    sourceType: 'module',
    ecmaVersion: 'latest',
    ecmaFeatures: { jsx: true },
    project: './tsconfig.json',
    tsconfigRootDir: __dirname,
  },

  settings: {
    'import/resolver': {
      typescript: {
        project: './tsconfig.json',
      },
    },
  },
  /**
   * 0 ~ 'off'
   * 1 ~ 'warn'
   * 2 ~ 'error'
   */
  rules: {
    'jsx-a11y/anchor-is-valid': 'off',
    'jsx-a11y/alt-text': 'off',
    'jsx-a11y/click-events-have-key-events': 'off',
    'jsx-a11y/no-static-element-interactions': 'off',
    'import/no-anonymous-default-export': 'off',
    '@typescript-eslint/no-unused-vars': 'off',
    'class-methods-use-this': 'off',
    '@typescript-eslint/no-unused-expressions': 'off',
    'no-unused-vars': 'off',
    'react/prop-types': 'off',
    'import/no-extraneous-dependencies': 'off',
    'perfectionist/sort-imports': 'off',
    'no-restricted-syntax': 'off',
    'react/button-has-type': 'off',
    'react/self-closing-comp': 'off',
    'react/no-danger': 'off',
    'arrow-body-style': 'off',
    'import/order': 'off',
    'unused-imports/no-unused-imports': 'off',
    '@typescript-eslint/no-unused-vars': 'off',
    'consistent-return': 'off',
    'react/no-unstable-nested-components': 'off',
    // general
    'no-alert': 0,
    camelcase: 0,
    'no-plusplus': 0,
    'no-multi-assign': 0,
    'no-console': 0,
    'no-unused-vars': 0,
    'no-nested-ternary': 0,
    'no-param-reassign': 0,
    'no-underscore-dangle': 0,
    'no-restricted-exports': 0,
    'no-promise-executor-return': 0,
    'import/prefer-default-export': 0,
    'prefer-destructuring': [1, { object: true, array: false }],
    // typescript
    '@typescript-eslint/naming-convention': 0,
    '@typescript-eslint/no-use-before-define': 0,
    '@typescript-eslint/consistent-type-exports': 1,
    '@typescript-eslint/consistent-type-imports': 'off',
    '@typescript-eslint/no-shadow': 0,
    // react
    'react/no-children-prop': 0,
    'react/react-in-jsx-scope': 0,
    'react/no-array-index-key': 0,
    'react/require-default-props': 0,
    'react/jsx-props-no-spreading': 0,
    'react/function-component-definition': 0,
    'react/jsx-no-duplicate-props': [1, { ignoreCase: false }],
    'react/jsx-no-useless-fragment': [1, { allowExpressions: true }],

    // jsx-a11y

    'jsx-a11y/control-has-associated-label': 0,
    // unused imports
    'unused-imports/no-unused-vars': [
      0,
      { vars: 'all', varsIgnorePattern: '^_', args: 'after-used', argsIgnorePattern: '^_' },
    ],
    // perfectionist
    'perfectionist/sort-exports': [1, { order: 'asc', type: 'line-length' }],
    'perfectionist/sort-named-imports': 'off',
    'perfectionist/sort-named-exports': [1, { order: 'asc', type: 'line-length' }],
  },
};
