'use strict';

var app = angular.module('app', [
  'sectionDirective',
  'stringFilter',
  'numberFilter',
  'ui.router',
  'ui.bootstrap',
  'ui.checkbox',
  'ui.sortable',
  'ui.select',
  'amazonCtrl',
  'indexCtrl',
  'cardCtrl',
  'deckCtrl',
  'mypageCtrl',
  'usersCtrl',
  'amazonService',
  'deckService',
  'userService',
  'angular-loading-bar',
  'ngAnimate',
  'simpleUiSelect',
  'trustFilter',
]);
app.config(['$httpProvider', '$locationProvider', '$stateProvider', '$urlRouterProvider', 'uiSelectConfig', function($httpProvider, $locationProvider, $stateProvider, $urlRouterProvider, uiSelectConfig) {
  $httpProvider.defaults.headers.common = {'X-Requested-With': 'XMLHttpRequest'};
  $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded; charset=UTF-8';
  uiSelectConfig.theme = 'selectize';
  $locationProvider.html5Mode({
    enabled: true,
    requireBase: false
  });
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
  $stateProvider.state('cardSearch', {
    url: '/search',
    views: {
      mainContent: {
        templateUrl: '/view/card/search.html',
        controller: 'cardSearchController'
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
  $stateProvider.state('deckDetail', {
    url: '/deck/:id',
    views: {
      mainContent: {
        templateUrl: '/view/deck/detail.html',
        controller: 'deckDetailController'
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
  $stateProvider.state('userDetail', {
    url: '/users/:userId',
    views: {
      mainContent: {
        templateUrl: '/view/users/detail.html',
        controller: 'userDetailController'
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
  $stateProvider.state('admin', {
    url: '/admin',
    views: {
      mainContent: {
        templateUrl: '/admin/amazon/index.html',
        controller: 'amazonController'
      }
    }
  });
}]);