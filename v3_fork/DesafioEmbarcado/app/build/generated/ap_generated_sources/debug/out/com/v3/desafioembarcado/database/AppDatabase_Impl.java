package com.v3.desafioembarcado.database;

import androidx.annotation.NonNull;
import androidx.room.DatabaseConfiguration;
import androidx.room.InvalidationTracker;
import androidx.room.RoomOpenHelper;
import androidx.room.RoomOpenHelper.Delegate;
import androidx.room.RoomOpenHelper.ValidationResult;
import androidx.room.migration.AutoMigrationSpec;
import androidx.room.migration.Migration;
import androidx.room.util.DBUtil;
import androidx.room.util.TableInfo;
import androidx.room.util.TableInfo.Column;
import androidx.room.util.TableInfo.ForeignKey;
import androidx.room.util.TableInfo.Index;
import androidx.sqlite.db.SupportSQLiteDatabase;
import androidx.sqlite.db.SupportSQLiteOpenHelper;
import androidx.sqlite.db.SupportSQLiteOpenHelper.Callback;
import androidx.sqlite.db.SupportSQLiteOpenHelper.Configuration;
import com.v3.desafioembarcado.database.dao.GPSDao;
import com.v3.desafioembarcado.database.dao.GPSDao_Impl;
import com.v3.desafioembarcado.database.dao.GyroscopeDao;
import com.v3.desafioembarcado.database.dao.GyroscopeDao_Impl;
import com.v3.desafioembarcado.database.dao.PhotoDao;
import com.v3.desafioembarcado.database.dao.PhotoDao_Impl;
import java.lang.Class;
import java.lang.Override;
import java.lang.String;
import java.lang.SuppressWarnings;
import java.util.Arrays;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;

@SuppressWarnings({"unchecked", "deprecation"})
public final class AppDatabase_Impl extends AppDatabase {
  private volatile GyroscopeDao _gyroscopeDao;

  private volatile GPSDao _gPSDao;

  private volatile PhotoDao _photoDao;

  @Override
  protected SupportSQLiteOpenHelper createOpenHelper(DatabaseConfiguration configuration) {
    final SupportSQLiteOpenHelper.Callback _openCallback = new RoomOpenHelper(configuration, new RoomOpenHelper.Delegate(1) {
      @Override
      public void createAllTables(SupportSQLiteDatabase _db) {
        _db.execSQL("CREATE TABLE IF NOT EXISTS `gyroscope_db` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `x` REAL NOT NULL, `y` REAL NOT NULL, `z` REAL NOT NULL, `timestamp` INTEGER NOT NULL)");
        _db.execSQL("CREATE TABLE IF NOT EXISTS `gps_db` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `latitude` REAL NOT NULL, `longitude` REAL NOT NULL, `timestamp` INTEGER NOT NULL)");
        _db.execSQL("CREATE TABLE IF NOT EXISTS `photo_db` (`id` INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, `filePath` TEXT, `timestamp` INTEGER NOT NULL)");
        _db.execSQL("CREATE TABLE IF NOT EXISTS room_master_table (id INTEGER PRIMARY KEY,identity_hash TEXT)");
        _db.execSQL("INSERT OR REPLACE INTO room_master_table (id,identity_hash) VALUES(42, 'e905a05cc251ddce86043c3dbe7cfcea')");
      }

      @Override
      public void dropAllTables(SupportSQLiteDatabase _db) {
        _db.execSQL("DROP TABLE IF EXISTS `gyroscope_db`");
        _db.execSQL("DROP TABLE IF EXISTS `gps_db`");
        _db.execSQL("DROP TABLE IF EXISTS `photo_db`");
        if (mCallbacks != null) {
          for (int _i = 0, _size = mCallbacks.size(); _i < _size; _i++) {
            mCallbacks.get(_i).onDestructiveMigration(_db);
          }
        }
      }

      @Override
      public void onCreate(SupportSQLiteDatabase _db) {
        if (mCallbacks != null) {
          for (int _i = 0, _size = mCallbacks.size(); _i < _size; _i++) {
            mCallbacks.get(_i).onCreate(_db);
          }
        }
      }

      @Override
      public void onOpen(SupportSQLiteDatabase _db) {
        mDatabase = _db;
        internalInitInvalidationTracker(_db);
        if (mCallbacks != null) {
          for (int _i = 0, _size = mCallbacks.size(); _i < _size; _i++) {
            mCallbacks.get(_i).onOpen(_db);
          }
        }
      }

      @Override
      public void onPreMigrate(SupportSQLiteDatabase _db) {
        DBUtil.dropFtsSyncTriggers(_db);
      }

      @Override
      public void onPostMigrate(SupportSQLiteDatabase _db) {
      }

      @Override
      public RoomOpenHelper.ValidationResult onValidateSchema(SupportSQLiteDatabase _db) {
        final HashMap<String, TableInfo.Column> _columnsGyroscopeDb = new HashMap<String, TableInfo.Column>(5);
        _columnsGyroscopeDb.put("id", new TableInfo.Column("id", "INTEGER", true, 1, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsGyroscopeDb.put("x", new TableInfo.Column("x", "REAL", true, 0, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsGyroscopeDb.put("y", new TableInfo.Column("y", "REAL", true, 0, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsGyroscopeDb.put("z", new TableInfo.Column("z", "REAL", true, 0, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsGyroscopeDb.put("timestamp", new TableInfo.Column("timestamp", "INTEGER", true, 0, null, TableInfo.CREATED_FROM_ENTITY));
        final HashSet<TableInfo.ForeignKey> _foreignKeysGyroscopeDb = new HashSet<TableInfo.ForeignKey>(0);
        final HashSet<TableInfo.Index> _indicesGyroscopeDb = new HashSet<TableInfo.Index>(0);
        final TableInfo _infoGyroscopeDb = new TableInfo("gyroscope_db", _columnsGyroscopeDb, _foreignKeysGyroscopeDb, _indicesGyroscopeDb);
        final TableInfo _existingGyroscopeDb = TableInfo.read(_db, "gyroscope_db");
        if (! _infoGyroscopeDb.equals(_existingGyroscopeDb)) {
          return new RoomOpenHelper.ValidationResult(false, "gyroscope_db(com.v3.desafioembarcado.database.entities.Gyroscope).\n"
                  + " Expected:\n" + _infoGyroscopeDb + "\n"
                  + " Found:\n" + _existingGyroscopeDb);
        }
        final HashMap<String, TableInfo.Column> _columnsGpsDb = new HashMap<String, TableInfo.Column>(4);
        _columnsGpsDb.put("id", new TableInfo.Column("id", "INTEGER", true, 1, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsGpsDb.put("latitude", new TableInfo.Column("latitude", "REAL", true, 0, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsGpsDb.put("longitude", new TableInfo.Column("longitude", "REAL", true, 0, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsGpsDb.put("timestamp", new TableInfo.Column("timestamp", "INTEGER", true, 0, null, TableInfo.CREATED_FROM_ENTITY));
        final HashSet<TableInfo.ForeignKey> _foreignKeysGpsDb = new HashSet<TableInfo.ForeignKey>(0);
        final HashSet<TableInfo.Index> _indicesGpsDb = new HashSet<TableInfo.Index>(0);
        final TableInfo _infoGpsDb = new TableInfo("gps_db", _columnsGpsDb, _foreignKeysGpsDb, _indicesGpsDb);
        final TableInfo _existingGpsDb = TableInfo.read(_db, "gps_db");
        if (! _infoGpsDb.equals(_existingGpsDb)) {
          return new RoomOpenHelper.ValidationResult(false, "gps_db(com.v3.desafioembarcado.database.entities.GPS).\n"
                  + " Expected:\n" + _infoGpsDb + "\n"
                  + " Found:\n" + _existingGpsDb);
        }
        final HashMap<String, TableInfo.Column> _columnsPhotoDb = new HashMap<String, TableInfo.Column>(3);
        _columnsPhotoDb.put("id", new TableInfo.Column("id", "INTEGER", true, 1, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsPhotoDb.put("filePath", new TableInfo.Column("filePath", "TEXT", false, 0, null, TableInfo.CREATED_FROM_ENTITY));
        _columnsPhotoDb.put("timestamp", new TableInfo.Column("timestamp", "INTEGER", true, 0, null, TableInfo.CREATED_FROM_ENTITY));
        final HashSet<TableInfo.ForeignKey> _foreignKeysPhotoDb = new HashSet<TableInfo.ForeignKey>(0);
        final HashSet<TableInfo.Index> _indicesPhotoDb = new HashSet<TableInfo.Index>(0);
        final TableInfo _infoPhotoDb = new TableInfo("photo_db", _columnsPhotoDb, _foreignKeysPhotoDb, _indicesPhotoDb);
        final TableInfo _existingPhotoDb = TableInfo.read(_db, "photo_db");
        if (! _infoPhotoDb.equals(_existingPhotoDb)) {
          return new RoomOpenHelper.ValidationResult(false, "photo_db(com.v3.desafioembarcado.database.entities.Photo).\n"
                  + " Expected:\n" + _infoPhotoDb + "\n"
                  + " Found:\n" + _existingPhotoDb);
        }
        return new RoomOpenHelper.ValidationResult(true, null);
      }
    }, "e905a05cc251ddce86043c3dbe7cfcea", "2a1fca5a9466b291437fece1e9ee0ef1");
    final SupportSQLiteOpenHelper.Configuration _sqliteConfig = SupportSQLiteOpenHelper.Configuration.builder(configuration.context)
        .name(configuration.name)
        .callback(_openCallback)
        .build();
    final SupportSQLiteOpenHelper _helper = configuration.sqliteOpenHelperFactory.create(_sqliteConfig);
    return _helper;
  }

  @Override
  protected InvalidationTracker createInvalidationTracker() {
    final HashMap<String, String> _shadowTablesMap = new HashMap<String, String>(0);
    HashMap<String, Set<String>> _viewTables = new HashMap<String, Set<String>>(0);
    return new InvalidationTracker(this, _shadowTablesMap, _viewTables, "gyroscope_db","gps_db","photo_db");
  }

  @Override
  public void clearAllTables() {
    super.assertNotMainThread();
    final SupportSQLiteDatabase _db = super.getOpenHelper().getWritableDatabase();
    try {
      super.beginTransaction();
      _db.execSQL("DELETE FROM `gyroscope_db`");
      _db.execSQL("DELETE FROM `gps_db`");
      _db.execSQL("DELETE FROM `photo_db`");
      super.setTransactionSuccessful();
    } finally {
      super.endTransaction();
      _db.query("PRAGMA wal_checkpoint(FULL)").close();
      if (!_db.inTransaction()) {
        _db.execSQL("VACUUM");
      }
    }
  }

  @Override
  protected Map<Class<?>, List<Class<?>>> getRequiredTypeConverters() {
    final HashMap<Class<?>, List<Class<?>>> _typeConvertersMap = new HashMap<Class<?>, List<Class<?>>>();
    _typeConvertersMap.put(GyroscopeDao.class, GyroscopeDao_Impl.getRequiredConverters());
    _typeConvertersMap.put(GPSDao.class, GPSDao_Impl.getRequiredConverters());
    _typeConvertersMap.put(PhotoDao.class, PhotoDao_Impl.getRequiredConverters());
    return _typeConvertersMap;
  }

  @Override
  public Set<Class<? extends AutoMigrationSpec>> getRequiredAutoMigrationSpecs() {
    final HashSet<Class<? extends AutoMigrationSpec>> _autoMigrationSpecsSet = new HashSet<Class<? extends AutoMigrationSpec>>();
    return _autoMigrationSpecsSet;
  }

  @Override
  public List<Migration> getAutoMigrations(
      @NonNull Map<Class<? extends AutoMigrationSpec>, AutoMigrationSpec> autoMigrationSpecsMap) {
    return Arrays.asList();
  }

  @Override
  public GyroscopeDao gyroscopeDao() {
    if (_gyroscopeDao != null) {
      return _gyroscopeDao;
    } else {
      synchronized(this) {
        if(_gyroscopeDao == null) {
          _gyroscopeDao = new GyroscopeDao_Impl(this);
        }
        return _gyroscopeDao;
      }
    }
  }

  @Override
  public GPSDao gpsDao() {
    if (_gPSDao != null) {
      return _gPSDao;
    } else {
      synchronized(this) {
        if(_gPSDao == null) {
          _gPSDao = new GPSDao_Impl(this);
        }
        return _gPSDao;
      }
    }
  }

  @Override
  public PhotoDao photoDao() {
    if (_photoDao != null) {
      return _photoDao;
    } else {
      synchronized(this) {
        if(_photoDao == null) {
          _photoDao = new PhotoDao_Impl(this);
        }
        return _photoDao;
      }
    }
  }
}
