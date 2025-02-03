package com.v3.desafioembarcado.database.dao;

import android.database.Cursor;
import androidx.room.EntityInsertionAdapter;
import androidx.room.RoomDatabase;
import androidx.room.RoomSQLiteQuery;
import androidx.room.SharedSQLiteStatement;
import androidx.room.util.CursorUtil;
import androidx.room.util.DBUtil;
import androidx.sqlite.db.SupportSQLiteStatement;
import com.v3.desafioembarcado.database.entities.Gyroscope;
import java.lang.Class;
import java.lang.Override;
import java.lang.String;
import java.lang.SuppressWarnings;
import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

@SuppressWarnings({"unchecked", "deprecation"})
public final class GyroscopeDao_Impl implements GyroscopeDao {
  private final RoomDatabase __db;

  private final EntityInsertionAdapter<Gyroscope> __insertionAdapterOfGyroscope;

  private final SharedSQLiteStatement __preparedStmtOfDeleteAll;

  private final SharedSQLiteStatement __preparedStmtOfDeleteBY;

  public GyroscopeDao_Impl(RoomDatabase __db) {
    this.__db = __db;
    this.__insertionAdapterOfGyroscope = new EntityInsertionAdapter<Gyroscope>(__db) {
      @Override
      public String createQuery() {
        return "INSERT OR ABORT INTO `gyroscope_db` (`id`,`x`,`y`,`z`,`timestamp`) VALUES (nullif(?, 0),?,?,?,?)";
      }

      @Override
      public void bind(SupportSQLiteStatement stmt, Gyroscope value) {
        stmt.bindLong(1, value.getId());
        stmt.bindDouble(2, value.getX());
        stmt.bindDouble(3, value.getY());
        stmt.bindDouble(4, value.getZ());
        stmt.bindLong(5, value.getTimestamp());
      }
    };
    this.__preparedStmtOfDeleteAll = new SharedSQLiteStatement(__db) {
      @Override
      public String createQuery() {
        final String _query = "DELETE FROM gyroscope_db";
        return _query;
      }
    };
    this.__preparedStmtOfDeleteBY = new SharedSQLiteStatement(__db) {
      @Override
      public String createQuery() {
        final String _query = "DELETE FROM gyroscope_db WHERE id = ?";
        return _query;
      }
    };
  }

  @Override
  public void insert(final Gyroscope gyroscope) {
    __db.assertNotSuspendingTransaction();
    __db.beginTransaction();
    try {
      __insertionAdapterOfGyroscope.insert(gyroscope);
      __db.setTransactionSuccessful();
    } finally {
      __db.endTransaction();
    }
  }

  @Override
  public void deleteAll() {
    __db.assertNotSuspendingTransaction();
    final SupportSQLiteStatement _stmt = __preparedStmtOfDeleteAll.acquire();
    __db.beginTransaction();
    try {
      _stmt.executeUpdateDelete();
      __db.setTransactionSuccessful();
    } finally {
      __db.endTransaction();
      __preparedStmtOfDeleteAll.release(_stmt);
    }
  }

  @Override
  public void deleteBY(final int id) {
    __db.assertNotSuspendingTransaction();
    final SupportSQLiteStatement _stmt = __preparedStmtOfDeleteBY.acquire();
    int _argIndex = 1;
    _stmt.bindLong(_argIndex, id);
    __db.beginTransaction();
    try {
      _stmt.executeUpdateDelete();
      __db.setTransactionSuccessful();
    } finally {
      __db.endTransaction();
      __preparedStmtOfDeleteBY.release(_stmt);
    }
  }

  @Override
  public List<Gyroscope> getAll() {
    final String _sql = "SELECT * FROM gyroscope_db ORDER BY timestamp DESC";
    final RoomSQLiteQuery _statement = RoomSQLiteQuery.acquire(_sql, 0);
    __db.assertNotSuspendingTransaction();
    final Cursor _cursor = DBUtil.query(__db, _statement, false, null);
    try {
      final int _cursorIndexOfId = CursorUtil.getColumnIndexOrThrow(_cursor, "id");
      final int _cursorIndexOfX = CursorUtil.getColumnIndexOrThrow(_cursor, "x");
      final int _cursorIndexOfY = CursorUtil.getColumnIndexOrThrow(_cursor, "y");
      final int _cursorIndexOfZ = CursorUtil.getColumnIndexOrThrow(_cursor, "z");
      final int _cursorIndexOfTimestamp = CursorUtil.getColumnIndexOrThrow(_cursor, "timestamp");
      final List<Gyroscope> _result = new ArrayList<Gyroscope>(_cursor.getCount());
      while(_cursor.moveToNext()) {
        final Gyroscope _item;
        final int _tmpId;
        _tmpId = _cursor.getInt(_cursorIndexOfId);
        final float _tmpX;
        _tmpX = _cursor.getFloat(_cursorIndexOfX);
        final float _tmpY;
        _tmpY = _cursor.getFloat(_cursorIndexOfY);
        final float _tmpZ;
        _tmpZ = _cursor.getFloat(_cursorIndexOfZ);
        final long _tmpTimestamp;
        _tmpTimestamp = _cursor.getLong(_cursorIndexOfTimestamp);
        _item = new Gyroscope(_tmpId,_tmpX,_tmpY,_tmpZ,_tmpTimestamp);
        _result.add(_item);
      }
      return _result;
    } finally {
      _cursor.close();
      _statement.release();
    }
  }

  public static List<Class<?>> getRequiredConverters() {
    return Collections.emptyList();
  }
}
