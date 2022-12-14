import { ValiQuery } from './types';
/**
 * Just a simple wrapper for a react component that is actually implementing the query editor.
 */
export class ValiAnnotationsQueryCtrl {
  static templateUrl = 'partials/annotations.editor.html';
  annotation: any;

  /** @ngInject */
  constructor() {
    this.annotation.target = this.annotation.target || {};
    this.onQueryChange = this.onQueryChange.bind(this);
  }

  onQueryChange(query: ValiQuery) {
    this.annotation.expr = query.expr;
    this.annotation.maxLines = query.maxLines;
    this.annotation.instant = query.instant;
  }
}
