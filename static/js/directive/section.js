var dir = angular.module('sectionDirective', [
  'cfp.loadingBar',
  'apiService',
  'amazonService',
]);

dir.directive('mainmenu', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/common/menu.html',
    controller: ['$scope', '$window', '$location', 'userService', 'cfpLoadingBar', function($scope, $window, $location, userService, cfpLoadingBar) {
      $scope.login = function() {
        $window.location.href = '/api/twitter/login';
      };
      $scope.loading = function() {
        cfpLoadingBar.start();
      };
      $scope.setHeaderClass = function() {
        if ($location.path() === '/') {
          $('header').addClass('top');
          $('header').removeClass('content');
        } else {
          $('header').addClass('content');
          $('header').removeClass('top');
        }
      };
      $scope.user = {};
      /**
       * セッションからログインユーザの情報をとる
       */
      userService.getLoginUser().then(function(data) {
        $scope.user = data; // 成功
      }, function(e) {
        $scope.user = undefined;
      });
    }]
  };
});

dir.directive('navbar', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/navbar.html'
  };
});

dir.directive('headerwrap', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/headerwrap.html'
  };
});

dir.directive('feature', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/feature.html'
  };
});

dir.directive('portfolio', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/portfolio.html'
  };
});

dir.directive('services', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/services.html'
  };
});

dir.directive('testimonials', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/testimonials.html'
  };
});

dir.directive('news', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/news.html'
  };
});

dir.directive('team', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/team.html'
  };
});

dir.directive('contact', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/contact.html'
  };
});

dir.directive('footer', function() {
  return {
    restrict: 'E',
    replace: false,
    templateUrl: '/view/common/footer.html'
  };
});

dir.directive('amazon', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/common/amazon.html',
    controller: ['$scope', 'amazonService', function($scope, amazonService) {
      $scope.amazonList = [];
      amazonService.search().then(function(data) {
        $scope.amazonList = data.slice(0, 6);
      });
    }]
  };
});

dir.directive('cardsearchform', ['$rootScope', function($rootScope) {
  return {
    restrict: 'E',
    replace: false,
    templateUrl: '/view/common/cardSearchForm.html',
    link: ['$scope', 'element', function($scope, element) {
    }],
    controller: ['$scope', '$rootScope', '$state', 'cardService', function($scope, $rootScope, $state, cardService) {
      if ($rootScope.searchCondition) {
        $scope.form = $rootScope.searchCondition;
      }
      if (!$scope.form) {
        $scope.form = { isDetail: false };
      }
      $scope.categories = ['ルリグ', 'アーツ', 'シグニ', 'スペル'];
      $scope.realities = ['LR', 'LC', 'SR', 'R', 'C', 'ST', 'PR'];
      $scope.levels = [0, 1, 2, 3, 4];
      $scope.powers = [1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 11000, 12000, 13000, 14000, 15000];
      $scope.costs = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];
      $scope.toggleDetail = function() {
        $scope.form.isDetail = !$scope.form.isDetail;
      };
      $scope.search = function() {
        $rootScope.searchCondition = angular.copy($scope.form);
        if ($state.current.name !== "cardSearch") {
          $state.go("cardSearch");
        } else {
          $scope.cardSearch();
        }
      };
      $scope.reset = function() {
        var isDetail = angular.copy($scope.form.isDetail);
        $scope.form = { isDetail: isDetail };
      };

      var init = function() {
        if (!$rootScope.illustrators) {
          cardService.getIllustrator().then(function(data) {
            $rootScope.illustrators = data;
          });
        }
        if (!$rootScope.constraints) {
          cardService.getConstraint().then(function(data) {
            $rootScope.constraints = data;
          });
        }
        if (!$rootScope.products) {
          cardService.getProduct().then(function(data) {
            $rootScope.products = data;
          });
        }
        if (!$rootScope.types) {
          cardService.getType().then(function(data) {
            $rootScope.types = data;
          });
        }
      };
      init();
    }]
  };
}]);

dir.directive('copyright', function() {
  return {
    restrict: 'E',
    replace: true,
    scope: {
      name: '@'
    },
    template: '<small>Copyright &copy; {{year}} {{name}} All Rights Reserved.</small>',
    link: ['$scope', function($scope) {
      $scope.year = new Date().getFullYear();
    }]
  };
});