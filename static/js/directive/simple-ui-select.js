var dir = angular.module('simpleUiSelect', []);
dir.directive('simpleSelect', function() {
  return {
    restrict: 'E',
    replace: true,
    scope: {
      value: '=ngModel',
      style: '@',
      items: '='
    },
    template: '<div>'
            + '  <ui-select ng-model="value" style="style" ng-change="setValue($select)">'
            + '    <ui-select-match placeholder="">{{value}}</ui-select-match>'
            + '    <ui-select-choices repeat="item in items">{{item}}</ui-select-choices>'
            + '  </ui-select>'
            + '</div>',
    link: ['$scope', function($scope) {
      $scope.$watch('value', function(newValue, oldValue) {
        if (!newValue) {
          $scope.value = undefined;
        }
      });
    }],
    controller: ['$scope', function($scope) {
      $scope.setValue = function(select) {
        $scope.value = select.selected;
      };
    }]
  };
});