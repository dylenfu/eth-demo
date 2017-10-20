pragma solidity ^0.4.15;


contract School {

    event SemenEvent(
        address[] addresses
    );

    event BabyEvent(
        address[3] addresses
    );

    event ChildEvent(
        address[3][] addressList
    );

    event StudentEvent(
        address[2][] addressList
    );

    event MatesEvent(
        address[2][]    addressList,
        uint[7][]       uintArgsList
    );

    event ClassEvent(
        address[2][]    addressList,
        uint[7][]       uintArgsList,
        uint8[]         vList
    );

    event GradeEvent(
        address[2][]    addressList,
        uint[7][]       uintArgsList,
        uint8[]         vList,
        bytes32[]       rList
    );

    function setSemen(address[] addresses) {
        SemenEvent(addresses);
    }

    function setBaby(address[3] addresses) {
        BabyEvent(addresses);
    }

    function setChild(address[3][] addressList) {
        ChildEvent(addressList);
    }

    function setStudent(address[2][] addressList) {
        StudentEvent(addressList);
    }

    function setMates(address[2][] addressList, uint[7][] matesList) {
        MatesEvent(addressList, matesList);
    }

    function setClass(
        address[2][] addressList,
        uint[7][] matesList,
        uint8[]   vList
    ) {
        ClassEvent(addressList, matesList, vList);
    }

    function setGrade(
        address[2][] addressList,
        uint[7][] matesList,
        uint8[]   vList,
        bytes32[]   rList
    ) {
        GradeEvent(addressList, matesList, vList, rList);
    }
}
