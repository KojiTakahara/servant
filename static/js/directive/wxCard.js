var dir = angular.module('wxCardDirective', []);
dir.directive('cardlist', function() {
  return {
    restrict: 'E',
    templateUrl: '/view/common/cardlist.html'
  };
});
dir.directive('carddetail', function() {
  return {
    restrict: 'E',
    templateUrl: '/view/common/cardDetail.html'
  };
});