import angular from 'angular';

const coreModule = angular.module('plutono.core', ['ngRoute']);

// legacy modules
const angularModules = [
  coreModule,
  angular.module('plutono.controllers', []),
  angular.module('plutono.directives', []),
  angular.module('plutono.factories', []),
  angular.module('plutono.services', []),
  angular.module('plutono.filters', []),
  angular.module('plutono.routes', []),
];

export { angularModules, coreModule };

export default coreModule;
