// Libraries
import React, { ReactNode } from 'react';

import {
  SlatePrism,
  TypeaheadOutput,
  SuggestionsState,
  QueryField,
  TypeaheadInput,
  BracesPlugin,
  DOMUtil,
  Icon,
} from '@credativ/plutono-ui';

// Utils & Services
// dom also includes Element polyfills
import { Plugin, Node } from 'slate';
import { Editor } from 'slate-react';
import { ValiLabelBrowser } from './ValiLabelBrowser';

// Types
import { ExploreQueryFieldProps, AbsoluteTimeRange } from '@credativ/plutono-data';
import { ValiQuery, ValiOptions } from '../types';
import { LanguageMap, languages as prismLanguages } from 'prismjs';
import ValiLanguageProvider, { ValiHistoryItem } from '../language_provider';
import ValiDatasource from '../datasource';
import ValiOptionFields from './ValiOptionFields';

function getChooserText(hasSyntax: boolean, hasLogLabels: boolean) {
  if (!hasSyntax) {
    return 'Loading labels...';
  }
  if (!hasLogLabels) {
    return '(No logs found)';
  }
  return 'Log browser';
}

function willApplySuggestion(suggestion: string, { typeaheadContext, typeaheadText }: SuggestionsState): string {
  // Modify suggestion based on context
  switch (typeaheadContext) {
    case 'context-labels': {
      const nextChar = DOMUtil.getNextCharacter();
      if (!nextChar || nextChar === '}' || nextChar === ',') {
        suggestion += '=';
      }
      break;
    }

    case 'context-label-values': {
      // Always add quotes and remove existing ones instead
      if (!typeaheadText.match(/^(!?=~?"|")/)) {
        suggestion = `"${suggestion}`;
      }
      if (DOMUtil.getNextCharacter() !== '"') {
        suggestion = `${suggestion}"`;
      }
      break;
    }

    default:
  }
  return suggestion;
}

export interface ValiQueryFieldFormProps extends ExploreQueryFieldProps<ValiDatasource, ValiQuery, ValiOptions> {
  history: ValiHistoryItem[];
  absoluteRange: AbsoluteTimeRange;
  ExtraFieldElement?: ReactNode;
  runOnBlur?: boolean;
}

interface ValiQueryFieldFormState {
  labelsLoaded: boolean;
  labelBrowserVisible: boolean;
}

export class ValiQueryFieldForm extends React.PureComponent<ValiQueryFieldFormProps, ValiQueryFieldFormState> {
  plugins: Array<Plugin<Editor>>;

  constructor(props: ValiQueryFieldFormProps) {
    super(props);

    this.state = { labelsLoaded: false, labelBrowserVisible: false };

    this.plugins = [
      BracesPlugin(),
      SlatePrism(
        {
          onlyIn: (node: Node) => node.object === 'block' && node.type === 'code_block',
          getSyntax: (node: Node) => 'logql',
        },
        { ...(prismLanguages as LanguageMap), logql: this.props.datasource.languageProvider.getSyntax() }
      ),
    ];
  }

  async componentDidUpdate() {
    await this.props.datasource.languageProvider.start();
    this.setState({ labelsLoaded: true });
  }

  onChangeLogLabels = (selector: string) => {
    this.onChangeQuery(selector, true);
    this.setState({ labelBrowserVisible: false });
  };

  onChangeQuery = (value: string, override?: boolean) => {
    // Send text change to parent
    const { query, onChange, onRunQuery } = this.props;
    if (onChange) {
      const nextQuery = { ...query, expr: value };
      onChange(nextQuery);

      if (override && onRunQuery) {
        onRunQuery();
      }
    }
  };

  onClickChooserButton = () => {
    this.setState((state) => ({ labelBrowserVisible: !state.labelBrowserVisible }));
  };

  onTypeahead = async (typeahead: TypeaheadInput): Promise<TypeaheadOutput> => {
    const { datasource } = this.props;

    if (!datasource.languageProvider) {
      return { suggestions: [] };
    }

    const valiLanguageProvider = datasource.languageProvider as ValiLanguageProvider;
    const { history } = this.props;
    const { prefix, text, value, wrapperClasses, labelKey } = typeahead;

    const result = await valiLanguageProvider.provideCompletionItems(
      { text, value, prefix, wrapperClasses, labelKey },
      { history }
    );
    return result;
  };

  render() {
    const { ExtraFieldElement, query, datasource, runOnBlur } = this.props;
    const { labelsLoaded, labelBrowserVisible } = this.state;
    const valiLanguageProvider = datasource.languageProvider as ValiLanguageProvider;
    const cleanText = datasource.languageProvider ? valiLanguageProvider.cleanText : undefined;
    const hasLogLabels = valiLanguageProvider.getLabelKeys().length > 0;
    const chooserText = getChooserText(labelsLoaded, hasLogLabels);
    const buttonDisabled = !(labelsLoaded && hasLogLabels);

    return (
      <>
        <div className="gf-form-inline gf-form-inline--xs-view-flex-column flex-grow-1">
          <button
            className="gf-form-label query-keyword pointer"
            onClick={this.onClickChooserButton}
            disabled={buttonDisabled}
          >
            {chooserText}
            <Icon name={labelBrowserVisible ? 'angle-down' : 'angle-right'} />
          </button>
          <div className="gf-form gf-form--grow flex-shrink-1 min-width-15">
            <QueryField
              additionalPlugins={this.plugins}
              cleanText={cleanText}
              query={query.expr}
              onTypeahead={this.onTypeahead}
              onWillApplySuggestion={willApplySuggestion}
              onChange={this.onChangeQuery}
              onBlur={this.props.onBlur}
              onRunQuery={this.props.onRunQuery}
              placeholder="Enter a Vali query (run with Shift+Enter)"
              portalOrigin="vali"
            />
          </div>
        </div>
        {labelBrowserVisible && (
          <div className="gf-form">
            <ValiLabelBrowser languageProvider={valiLanguageProvider} onChange={this.onChangeLogLabels} />
          </div>
        )}
        <ValiOptionFields
          queryType={query.instant ? 'instant' : 'range'}
          lineLimitValue={query?.maxLines?.toString() || ''}
          query={query}
          onRunQuery={this.props.onRunQuery}
          onChange={this.props.onChange}
          runOnBlur={runOnBlur}
        />
        {ExtraFieldElement}
      </>
    );
  }
}
