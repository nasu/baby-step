pragma solidity ^0.4.15;
contract ReserverRoom {
    // data
    address owner;
    string[] public rooms = ["roomA","roomB","roomC"];
    struct ReservationData {
        uint reserveStart;
        uint reserveEnd;
        address user;
    }
    struct UsageData {
        uint usageTime;
        address user;
    }
    struct RoomInfo {
        ReservationData[] reservationDB;
        UsageData[] usageLog;
        bool isValue;
    }
    mapping(string => RoomInfo) roomData;

    // modifier
    modifier isExist(string _room) {
        require (roomData[_room].isValue == true);
        _;
    }
    modifier isOwner() {
        require (msg.sender == owner);
        _;
    }

    // event
    event ReserveLog(uint _now, string room, uint _start, uint _end, address _user);
    event RoomUsageLog(uint _time, address _user);

    // constructor
    constructor() public {
        for (uint i = 0; i < rooms.length; i++) {
            /*
            ReservationData[] memory db = new ReservationData[](0);
            UsageData[] memory log = new UsageData[](0);
            roomData[rooms[i]] = RoomInfo({
                reservationDB: db,
                usageLog: log,
                isValue: true
            });
            */
            roomData[rooms[i]].isValue = true;
        }
        owner = msg.sender;
    }

    // function
    function offer(uint _start, uint _end, string _room)
        isExist(_room)
        public returns(bool)
    {
        if (now > _start) return false;
        if (isReserved(_start, _end, _room) == true) return false;
        roomData[_room].reservationDB.push(
            ReservationData({
                reserveStart: _start,
                reserveEnd: _end,
                user: msg.sender
            })
        );
        emit ReserveLog(now, _room, _start, _end, msg.sender);
        return true;
    }

    function isReserved(uint _start, uint _end, string _room)
        internal view returns(bool)
    {
        for (uint i = 0; i < roomData[_room].reservationDB.length; i++) {
            ReservationData storage r = roomData[_room].reservationDB[i];
            if (_end <= r.reserveStart) continue;
            if (r.reserveEnd <= _start) continue;
            return true;
        }
        return false;
    }

    function use(string _room)
        isExist(_room)
        public returns(bool)
    {
        uint time = now;
        if (isAvailable(time, _room) == false) return false;
        roomData[_room].usageLog.push(
            UsageData({
                usageTime: time,
                user: msg.sender
            })
        );
        emit RoomUsageLog(time, msg.sender);
        return true;
    }

    function isAvailable(uint _time, string _room)
        internal view returns(bool)
    {
        for (uint i = 0; i < roomData[_room].reservationDB.length; i++) {
            ReservationData storage r = roomData[_room].reservationDB[i];
            if (_time < r.reserveStart) continue;
            if (r.reserveEnd < _time) continue;
            if (msg.sender != r.user) continue;
            return true;
        }
        return false;
    }

    function getReservationDB(string _room)
        isExist(_room)
        public
    {
        for (uint i = 0; i < roomData[_room].reservationDB.length; i++) {
            ReservationData storage r = roomData[_room].reservationDB[i];
            emit ReserveLog(now, _room, r.reserveStart, r.reserveEnd, r.user);
        }
    }

    function getUsageLog(string _room)
        isExist(_room)
        public
    {
        for (uint i = 0; i < roomData[_room].usageLog.length; i++) {
            UsageData storage u = roomData[_room].usageLog[i];
            emit RoomUsageLog(u.usageTime, u.user);
        }
    }

    function cleanObsolete(string _room)
        isOwner isExist(_room)
        public returns (bool)
    {
        ReservationData[] storage olds = roomData[_room].reservationDB;
        ReservationData[] memory news = new ReservationData[](olds.length);
        uint newCount = 0;
        for (uint i = 0; i < olds.length; i++) {
            if (now <= olds[i].reserveEnd) continue;
            news[newCount++] = olds[i];
        }
        olds.length = 0;
        for (uint j = 0; j < newCount; j++) {
            //TODO: なぜこれで十分なのか. storage/memory の違いとは
            olds.push(news[j]);
        }
        return true;
    }
}