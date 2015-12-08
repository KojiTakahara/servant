'use strict';

var ctrl = angular.module('amazonCtrl', [
]);

ctrl.controller('amazonController', ['$scope', 'amazonService', function($scope, amazonService) {

  $scope.amazonList = [];

  $scope.create = function() {
    amazonService.create($scope.form).then(function(data) {
      alert("success");
    });
  };

  $scope.edit = function(amazon) {
    $scope.form = {
      name: amazon.Name,
      weight: amazon.Weight,
      enabled: amazon.Enabled,
      html: amazon.Html
    }
  };

  var init = function() {
    amazonService.search().then(function(data) {
      $scope.amazonList = data;
    });
  };
  init();

}]);
