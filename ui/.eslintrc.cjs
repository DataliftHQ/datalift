module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  settings: {
    'react': {
      'version': 'detect'
    },
    'import/extensions': ['.js', '.jsx', '.ts', '.tsx'],
    'import/resolver': {
      'node': {
        'extensions': ['.js', '.jsx', '.ts', '.tsx'],
        'paths': ['**/src', '**/dist']
      }
    },
    'import/parsers': {
      '@typescript-eslint/parser': ['.ts', '.tsx']
    }
  },
  extends: [
    'prettier',
    'eslint:recommended',
    'plugin:cypress/recommended',
    'plugin:import/typescript',
    'plugin:jest/recommended',
    'plugin:jest/style',
    'plugin:react/recommended',
    'plugin:jsx-a11y/recommended',
    'plugin:react-hooks/recommended',
    'plugin:prettier/recommended',
    'plugin:testing-library/react'
  ],
  parser: '@typescript-eslint/parser',
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
    ecmaFeatures: {
      jsx: true
    }
  },
  plugins: ['react-refresh', 'simple-import-sort'],
  rules: {
    'react-refresh/only-export-components': 'warn',
    'prettier/prettier': ['error', {}, { 'usePrettierrc': true }],
    'simple-import-sort/imports': 'error',
    'simple-import-sort/exports': 'error',
    'consistent-return': ['off'],
    'jest/expect-expect': [
      'error',
      {
        'assertFunctionNames': ['expect', 'cy']
      }
    ],
    'no-console': ['error'],
    'no-empty': ['off'],
    'no-nested-ternary': ['off'],
    'no-unused-expressions': ['error', { 'allowShortCircuit': true }],
    'react-hooks/exhaustive-deps': ['off'],
    'react/prop-types': ['off'],
    'react/require-default-props': ['off'],
    'react/jsx-props-no-spreading': [
      'off',
      {
        'exceptions': ['Wizard']
      }
    ],
    'jsx-a11y/accessible-emoji': 'off',
    'jsx-a11y/label-has-associated-control': [
      'error',
      {
        'required': {
          'some': ['nesting', 'id']
        }
      }
    ]
  },
  overrides: [
    {
      files: ['**/*.ts', '**/*.tsx'],
      parser: '@typescript-eslint/parser',
      plugins: ['@typescript-eslint'],
      rules: {
        'no-undef': ['off'],
        'no-unused-vars': ['off'],
        '@typescript-eslint/no-unused-vars': ['error', { 'args': 'none' }],
        'no-use-before-define': ['off'],
        '@typescript-eslint/no-use-before-define': ['error'],
        'no-shadow': ['off'],
        '@typescript-eslint/no-shadow': ['error']
      }
    }
  ]
}
