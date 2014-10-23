'use strict';

var app = angular.module('app', [
  'sectionDirective',
  'stringFilter',
  'numberFilter',
  'ui.router',
  'ui.bootstrap',
  'cardCtrl',
  'ngSelect',
  'angular-loading-bar',
  'ngAnimate'
]);
app.config(['$locationProvider', '$stateProvider', '$urlRouterProvider', function($locationProvider, $stateProvider, $urlRouterProvider) {
  $locationProvider.html5Mode(true);
  $urlRouterProvider.otherwise("");
  $stateProvider.state('top', {
    url: '/',
    views: {
      headerContent: { templateUrl: '/view/welcome.html' },
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
  /**
  $stateProvider.state('card.expansion.detail', {
    url: '/:no',
    views: {
      mainContent: {
        templateUrl: '/view/card/index.html',
        controller: 'cardController'
      }
    }
  });
  */
}]);