'use strict';

var app = angular.module('usersCtrl', [
  'apiService',
]);
app.controller('userDetailController', ['$scope', '$stateParams', '$filter', 'cardService', function($scope, $stateParams, $filter, cardService) {
  $scope.decks = [];
  $scope.alerts = [];
  $scope.closeAlert = function(index) {
    $scope.alerts.splice(index, 1);
  };
  var init = function() {
    var userId = $stateParams.userId;
    cardService.getUser(userId).then(function(data) {
      cardService.getTwitterUser(userId).then(function(res) {
        $scope.userInfo = res;
        $scope.userInfo.profile_image_url = $filter('mediumImage')($scope.userInfo.profile_image_url);
      });
      cardService.getDeckByUserId(userId, 'PUBLIC').then(function(res) {
        $scope.decks = res;
      });
    }, function() {
      $scope.alerts.push({ type: 'danger', msg: 'データが存在しません。' });
    });
  };
  init();
}]);
app.controller('usersController', ['$scope', '$window', function($scope, $window) {
}]);