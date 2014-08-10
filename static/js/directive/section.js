var dir = angular.module('sectionDirective', []);

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
    restrict: 'A',
    replace: true,
    templateUrl: '/view/footer.html'
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