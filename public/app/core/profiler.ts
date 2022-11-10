import { PlutonoRootScope } from 'app/routes/PlutonoCtrl';

export class Profiler {
  panelsRendered: number;
  enabled: boolean;
  $rootScope: PlutonoRootScope;
  window: any;

  init(config: any, $rootScope: PlutonoRootScope) {
    this.$rootScope = $rootScope;
    this.window = window;

    if (!this.enabled) {
      return;
    }
  }

  renderingCompleted() {
    // add render counter to root scope
    // used by image renderer to know when panel has rendered
    this.panelsRendered = (this.panelsRendered || 0) + 1;

    // this window variable is used by backend rendering tools to know
    // all panels have completed rendering
    this.window.panelsRendered = this.panelsRendered;
  }
}

const profiler = new Profiler();
export { profiler };
