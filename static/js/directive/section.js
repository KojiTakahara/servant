var dir = angular.module('sectionDirective', []);

dir.directive('mainmenu', function() {
  return {
    restrict: 'E',
    replace: true,
    templateUrl: '/view/common/menu.html',
    controller: function($scope, $location, userService) {
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
        console.log(data);
      }, function(e) {
        $scope.user = undefined;
        console.log(e); // ログインユーザが見当たらない
      });
    }
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
    replace: false,
    templateUrl: '/view/common/amazon.html'
  };
});

dir.directive('copyright', function() {
  return {
    restrict: 'E',
    replace: true,
    scope: {
      name: '@'
    },
    template: '<small>Copyright &copy; {{year}} {{name}} All Rights Reserved.</small>',
    link: function($scope) {
      $scope.year = new Date().getFullYear();
    }
  };
});