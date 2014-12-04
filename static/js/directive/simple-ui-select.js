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
    priority: 1,
    template: '<div>'
            + '  <ui-select ng-model="value" style="style" ng-change="setValue($select)">'
            + '    <ui-select-match placeholder="">{{value}}</ui-select-match>'
            + '    <ui-select-choices repeat="item in items">{{item}}</ui-select-choices>'
            + '  </ui-select>'
            + '</div>',
    link: function($scope) {
      $scope.$watch('value', function(newValue, oldValue) {
        if (!newValue) {
          console.log(1);
          $scope.value = undefined;
        }
      });
    },
    controller: function($scope) {
      $scope.setValue = function(select) {
        $scope.value = select.selected;
      };
    }
  };
});