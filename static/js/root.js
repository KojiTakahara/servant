'use strict';

var app = angular.module('app', [
  'sectionDirective',
  'stringFilter',
  'numberFilter',
  'ui.router',
  'ui.bootstrap',
  'indexCtrl',
  'cardCtrl',
  'userService',
  'ngSelect',
  'angular-loading-bar',
  'ngAnimate',
  'ui.select'
]);
app.config(['$locationProvider', '$stateProvider', '$urlRouterProvider', 'uiSelectConfig', function($locationProvider, $stateProvider, $urlRouterProvider, uiSelectConfig) {
  $locationProvider.html5Mode(true);
  $urlRouterProvider.otherwise("");
  uiSelectConfig.theme = 'selectize';
  $stateProvider.state('top', {
    url: '/',
    views: {
      headerContent: {
        templateUrl: '/view/welcome.html',
        controller: 'indexController'
      },
      mainContent: { templateUrl: '/view/top/menuDescription.html' }
    }
  });
  $stateProvider.state('card', {
    url: '/card',
    views: {
      mainContent: {
        templateUrl: '/view/card/index.html',
        controller: 'cardController'
      }
    }
  });
  $stateProvider.state('expansion', {
    url: '/card/:expansion',
    views: {
      mainContent: {
        templateUrl: '/view/card/expansion.html',
        controller: 'cardExController'
      }
    }
  });
  $stateProvider.state('cardDetail', {
    url: '/card/:expansion/:no',
    views: {
      mainContent: {
        templateUrl: '/view/card/detail.html',
        controller: 'cardDetailController'
      }
    }
  });
}]);