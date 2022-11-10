import { SelectableValue } from '@credativ/plutono-data';
import { PlutonoAnnotationType } from './types';

export const annotationTypes: Array<SelectableValue<PlutonoAnnotationType>> = [
  { text: 'Dashboard', value: PlutonoAnnotationType.Dashboard },
  { text: 'Tags', value: PlutonoAnnotationType.Tags },
];

export class PlutonoAnnotationsQueryCtrl {
  annotation: any;

  types = annotationTypes;

  constructor() {
    this.annotation.type = this.annotation.type || PlutonoAnnotationType.Tags;
    this.annotation.limit = this.annotation.limit || 100;
  }

  static templateUrl = 'partials/annotations.editor.html';
}
