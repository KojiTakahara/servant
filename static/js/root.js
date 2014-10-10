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
  var provider = $stateProvider;
  provider.state('top', {
    url: '/',
    views: {
      headerContent: { templateUrl: '/view/welcome.html' },
      mainContent: { templateUrl: '/view/top/menuDescription.html' }
    }
  });
  provider.state('card', {
    url: '/card',
    views: {
      mainContent: {
        templateUrl: '/view/card/index.html',
        controller: 'cardController'
      }
    }
  });
  provider.state('card.expansion', {
    url: '/:expansion',
    views: {
      mainContent: {
        templateUrl: '/view/card/index.html',
        controller: 'cardController'
      }
    }
  });
  provider.state('card.expansion.detail', {
    url: '/:no',
    views: {
      mainContent: {
        templateUrl: '/view/card/index.html',
        controller: 'cardController'
      }
    }
  });
}]);