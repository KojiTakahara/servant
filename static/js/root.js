'use strict';

var app = angular.module('app', [
  'sectionDirective',
  'stringFilter',
  'numberFilter',
  'ui.router',
  'ui.bootstrap',
  'indexCtrl',
  'cardCtrl',
  'deckCtrl',
  'mypageCtrl',
  'usersCtrl',
  'deckService',
  'userService',
  'selectize',
  'angular-loading-bar',
  'ngAnimate',
]);
app.config(['$locationProvider', '$stateProvider', '$urlRouterProvider', function($locationProvider, $stateProvider, $urlRouterProvider) {
  $locationProvider.html5Mode(true);
  $urlRouterProvider.otherwise("");
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
  $stateProvider.state('deck', {
    url: '/deck',
    views: {
      mainContent: {
        templateUrl: '/view/deck/index.html',
        controller: 'deckController'
      }
    }
  });
  $stateProvider.state('users', {
    url: '/users',
    views: {
      mainContent: {
        templateUrl: '/view/users/index.html',
        controller: 'usersController'
      }
    }
  });
  $stateProvider.state('mypage', {
    url: '/mypage',
    views: {
      mainContent: {
        templateUrl: '/view/mypage/index.html',
        controller: 'mypageController'
      }
    }
  });
  $stateProvider.state('mydeck', {
    url: '/mypage/deck/:id',
    views: {
      mainContent: {
        templateUrl: '/view/mypage/edit.html',
        controller: 'editDeckController'
      }
    }
  });
}]);